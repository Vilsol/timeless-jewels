import * as env from '$app/env';
import SyncWorker from './sync_worker?worker';
import * as Comlink from 'comlink';

function getWorker() {
  console.log('Creating sync worker');
  const theWorker = new SyncWorker();
  const obj = Comlink.wrap(theWorker);
  return { syncWorker: theWorker, syncWrap: obj };
}

export const { syncWorker, syncWrap } = env.browser ? getWorker() : { syncWorker: null, syncWrap: null };
