import { browser } from '$app/environment';
import { calculator } from './types';

// Tracing for the seed-preview path. Everything here is inert until enabled, so
// the instrumentation can stay in place without costing anything in normal use.
// Enable with ?perf=1, or __perf.on() from the console.

type Bucket = {
  count: number;
  total: number;
  max: number;
  last: number;
};

declare global {
  interface Window {
    __perf: {
      on: () => void;
      off: () => void;
      report: () => string;
      reset: () => void;
    };
  }
}

const buckets: Record<string, Bucket> = {};

let enabled = false;
let reportTimer: ReturnType<typeof setTimeout> | undefined;
let frameHandle: number | undefined;
let lastFrame = 0;
let observer: PerformanceObserver | undefined;
let syncing = false;

export const isEnabled = (): boolean => enabled;

const write = (label: string, ms: number) => {
  const bucket = buckets[label] || (buckets[label] = { count: 0, total: 0, max: 0, last: 0 });
  bucket.count++;
  bucket.total += ms;
  bucket.last = ms;
  if (ms > bucket.max) {
    bucket.max = ms;
  }
};

// Before initializeCrystalline() runs, the namespace is a proxy that throws on
// any read, so readiness cannot be feature-detected — it has to be attempted.
const whenReady = <T>(call: () => T): T | undefined => {
  try {
    return call();
  } catch {
    return undefined;
  }
};

const pad = (text: string, width: number) => text.padEnd(width);
const num = (value: number, width: number, digits = 1) => value.toFixed(digits).padStart(width);

export const report = (): string => {
  const labels = Object.keys(buckets).sort((a, b) => buckets[b].total - buckets[a].total);
  const width = Math.max(24, ...labels.map((l) => l.length));

  const lines = [
    `${pad('label', width)} ${'count'.padStart(7)} ${'total ms'.padStart(10)} ${'avg ms'.padStart(
      9
    )} ${'max ms'.padStart(9)}`,
    '-'.repeat(width + 39)
  ];

  labels.forEach((label) => {
    const b = buckets[label];
    const avg = b.count ? b.total / b.count : 0;
    lines.push(
      `${pad(label, width)} ${String(b.count).padStart(7)} ${num(b.total, 10)} ${num(avg, 9, 3)} ${num(b.max, 9, 3)}`
    );
  });

  const stats = whenReady(() => calculator.GetCalculateStats());
  if (!stats) {
    lines.push('');
    lines.push('go: unavailable — the wasm module was not up when the report ran');
  } else if (!stats.Calls) {
    lines.push('');
    lines.push('go: no calls recorded — tracking was not enabled on the Go side');
  }

  if (stats && stats.Calls) {
    const goTotal = stats.TotalNanos / 1e6;
    lines.push('');
    lines.push(
      `${pad('go: Calculate', width)} ${String(stats.Calls).padStart(7)} ${num(goTotal, 10)} ${num(
        goTotal / stats.Calls,
        9,
        3
      )} ${num(stats.MaxNanos / 1e6, 9, 3)}`
    );

    const js = buckets['wasm: Calculate'];
    if (js && js.total > goTotal) {
      const share = (1 - goTotal / js.total) * 100;
      lines.push(`${pad('go: boundary overhead', width)} ${num(js.total - goTotal, 18)} ms of the JS wall time`);
      lines.push(`${pad('go: boundary share', width)} ${num(share, 18)} %`);
    }
  }

  const text = lines.join('\n');
  console.log('%c[perf]\n' + text, 'font-family: monospace');

  return text;
};

// The auto-report fires once activity stops. Only work the user triggered
// postpones it — the frame loop records continuously and would hold it off
// forever.
const scheduleReport = () => {
  clearTimeout(reportTimer);
  reportTimer = setTimeout(() => report(), 1000);
};

export const record = (label: string, ms: number): void => {
  if (!enabled) {
    return;
  }

  write(label, ms);
  scheduleReport();
};

export const time = <T>(label: string, fn: () => T): T => {
  if (!enabled) {
    return fn();
  }

  const start = performance.now();
  try {
    return fn();
  } finally {
    record(label, performance.now() - start);
  }
};

// count records an occurrence with no duration of its own, for things that are
// interesting by frequency rather than by cost.
export const count = (label: string): void => record(label, 0);

// sinceInput measures from a user input to the frame that paints after it, which
// is the number that actually corresponds to "feels sluggish".
export const sinceInput = (label: string): void => {
  if (!enabled) {
    return;
  }

  const start = performance.now();
  requestAnimationFrame(() => {
    requestAnimationFrame(() => record(label, performance.now() - start));
  });
};

const frameLoop = () => {
  const now = performance.now();
  if (lastFrame) {
    const delta = now - lastFrame;
    // An idle tab throttles to a crawl and would otherwise dominate the numbers.
    if (delta < 1000) {
      write('frame delta', delta);
      if (delta > 50) {
        write('frames over 50ms', 0);
      }
    }
  }

  lastFrame = now;
  frameHandle = requestAnimationFrame(frameLoop);
};

// The wasm module comes up well after this module loads, so the Go-side toggle
// is retried until there is something to toggle.
const syncGoTracking = () => {
  let done = false;
  whenReady(() => {
    calculator.SetCalculateTracking(enabled);
    done = true;
  });

  if (done) {
    syncing = false;
    return;
  }

  syncing = true;
  setTimeout(syncGoTracking, 200);
};

export const reset = (): void => {
  Object.keys(buckets).forEach((key) => delete buckets[key]);
  whenReady(() => calculator.ResetCalculateStats());
  console.log('[perf] counters cleared');
};

export const setEnabled = (value: boolean): void => {
  enabled = value;

  if (!syncing) {
    syncGoTracking();
  }

  if (value && frameHandle === undefined) {
    lastFrame = 0;
    frameHandle = requestAnimationFrame(frameLoop);

    try {
      observer = new PerformanceObserver((list) => {
        list.getEntries().forEach((entry) => write('long task', entry.duration));
      });
      observer.observe({ entryTypes: ['longtask'] });
    } catch {
      // longtask is not observable everywhere; the frame loop still covers it.
    }
  }

  if (!value && frameHandle !== undefined) {
    cancelAnimationFrame(frameHandle);
    frameHandle = undefined;
    observer?.disconnect();
    observer = undefined;
  }

  console.log(`[perf] tracking ${value ? 'enabled' : 'disabled'}`);
};

const STORAGE_KEY = 'perf';

if (browser) {
  window.__perf = {
    on: () => {
      localStorage.setItem(STORAGE_KEY, '1');
      setEnabled(true);
    },
    off: () => {
      localStorage.removeItem(STORAGE_KEY);
      setEnabled(false);
    },
    report,
    reset
  };

  // Persisted, because ?perf=1 is lost the moment the page calls goto() to
  // rewrite the URL, and the landing page navigates before the tree is reached.
  if (new URLSearchParams(window.location.search).get('perf') === '1') {
    localStorage.setItem(STORAGE_KEY, '1');
  }

  if (localStorage.getItem(STORAGE_KEY) === '1') {
    setEnabled(true);
  } else {
    console.log('[perf] tracking is off — run __perf.on() to trace the seed path');
  }
}
