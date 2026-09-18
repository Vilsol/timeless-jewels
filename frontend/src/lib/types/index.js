/* eslint-disable */

const wrap = (name, fn) => {
  return (...args) => {
    const result = fn.call(undefined, ...args);
    if (globalThis.goInternalError) {
      const error = new Error(globalThis.goInternalError);
      globalThis.goInternalError = undefined;
      throw error;
    }
    return result;
  }
};

let initialized = false;

const pending = (name) => new Proxy({}, {
  get(target, property) {
    if (typeof property === 'symbol' || property === 'then') {
      return undefined;
    }
    if (initialized) {
      throw new Error('crystalline: this ' + name + ' was captured before initializeCrystalline() ran, so it is a stale copy. Read it from the module instead of destructuring it earlier, or move the import after initialisation.');
    }
    throw new Error('crystalline: ' + name + '.' + String(property) + ' was read before initializeCrystalline() ran. Start the Go wasm module, then call initializeCrystalline().');
  }
});

export let calculator = pending('calculator');
export let data = pending('data');

export const initializeCrystalline = () => {
  if (globalThis['go']?.['timeless-jewels'] === undefined) {
    throw new Error('crystalline: globalThis.go.timeless-jewels is not set. Start the Go wasm module before calling initializeCrystalline().');
  }

  const failedImports = globalThis['go']['timeless-jewels']['__crystalline']?.['importFailures'];
  if (failedImports?.length) {
    throw new Error('crystalline: Go could not reach what it imported: ' + failedImports.join('; '));
  }

  calculator = {
    Calculate: wrap('calculator.Calculate', globalThis['go']['timeless-jewels']['calculator']['Calculate']),
    GetCalculateStats: wrap('calculator.GetCalculateStats', globalThis['go']['timeless-jewels']['calculator']['GetCalculateStats']),
    ResetCalculateStats: wrap('calculator.ResetCalculateStats', globalThis['go']['timeless-jewels']['calculator']['ResetCalculateStats']),
    ReverseSearch: wrap('calculator.ReverseSearch', globalThis['go']['timeless-jewels']['calculator']['ReverseSearch']),
    SetCalculateTracking: wrap('calculator.SetCalculateTracking', globalThis['go']['timeless-jewels']['calculator']['SetCalculateTracking'])
  };
  data = {
    Conqueror: globalThis['go']['timeless-jewels']['data']['Conqueror'],
    GetAlternatePassiveAdditionByIndex: wrap('data.GetAlternatePassiveAdditionByIndex', globalThis['go']['timeless-jewels']['data']['GetAlternatePassiveAdditionByIndex']),
    GetAlternatePassiveSkillByIndex: wrap('data.GetAlternatePassiveSkillByIndex', globalThis['go']['timeless-jewels']['data']['GetAlternatePassiveSkillByIndex']),
    GetPassiveSkillByIndex: wrap('data.GetPassiveSkillByIndex', globalThis['go']['timeless-jewels']['data']['GetPassiveSkillByIndex']),
    GetStatByIndex: wrap('data.GetStatByIndex', globalThis['go']['timeless-jewels']['data']['GetStatByIndex']),
    JewelType: globalThis['go']['timeless-jewels']['data']['JewelType'],
    PassiveSkillAuraStatTranslationsJSON: globalThis['go']['timeless-jewels']['data']['PassiveSkillAuraStatTranslationsJSON'],
    PassiveSkillStatTranslationsJSON: globalThis['go']['timeless-jewels']['data']['PassiveSkillStatTranslationsJSON'],
    PassiveSkillType: globalThis['go']['timeless-jewels']['data']['PassiveSkillType'],
    PassiveSkills: globalThis['go']['timeless-jewels']['data']['PassiveSkills'],
    PossibleStats: globalThis['go']['timeless-jewels']['data']['PossibleStats'],
    SkillTree: globalThis['go']['timeless-jewels']['data']['SkillTree'],
    StatTranslationsJSON: globalThis['go']['timeless-jewels']['data']['StatTranslationsJSON'],
    TimelessJewelConquerors: globalThis['go']['timeless-jewels']['data']['TimelessJewelConquerors'],
    TimelessJewelSeedRanges: globalThis['go']['timeless-jewels']['data']['TimelessJewelSeedRanges'],
    TimelessJewels: globalThis['go']['timeless-jewels']['data']['TimelessJewels'],
    TreeToPassive: globalThis['go']['timeless-jewels']['data']['TreeToPassive']
  };

  initialized = true;
};

export const boot = async (wasm) => {
  if (globalThis['Go'] === undefined) {
    throw new Error('crystalline: the Go runtime shim is missing. Load wasm_exec.js from your Go toolchain before calling boot().');
  }

  const runtime = new globalThis['Go']();

  let source = wasm;

  if (!(wasm instanceof ArrayBuffer) && !ArrayBuffer.isView(wasm)) {
    const response = await fetch(wasm);

    if (!response.ok) {
      throw new Error('crystalline: could not fetch ' + wasm + ': ' + response.status);
    }

    source = await response.arrayBuffer();
  }

  const { instance } = await WebAssembly.instantiate(source, runtime.importObject);

  // Not awaited: the Go program parks, so this never settles.
  runtime.run(instance);

  initializeCrystalline();

  return { calculator, data };
};