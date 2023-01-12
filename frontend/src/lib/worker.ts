import { browser } from '$app/environment';
import SyncWorker from './sync_worker?worker';
import * as Comlink from 'comlink';
import type { WorkerType } from './sync_worker';

function getWorker() {
  console.log('Creating sync worker');
  const theWorker = new SyncWorker();
  const obj = Comlink.wrap<WorkerType>(theWorker);
  return { syncWorker: theWorker, syncWrap: obj };
}

export const { syncWorker, syncWrap } = browser ? getWorker() : { syncWorker: null, syncWrap: null };
