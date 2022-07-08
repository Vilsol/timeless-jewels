import { expose } from 'comlink';
import '../wasm_exec.js';

const obj = {
  boot(data: ArrayBuffer) {
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    const go = new Go();
    WebAssembly.instantiate(data, go.importObject).then((result) => {
      go.run(result.instance);
    });
  },
  search(args: unknown[], callback: (seed: number) => void) {
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    return ReverseSearch(...args, callback);
  }
};

expose(obj);
