import{S as R,i as C,s as S,l as U,g as z,o as k,p as D,q as T,d as g,n as A,F as N,G as P,H as O,I as V,e as b,t as j,k as F,c as y,a as x,h as I,m as M,b as v,J as p,E as _}from"../chunks/index-40c9b43b.js";import{a as G}from"../chunks/paths-396f020f.js";(()=>{const n=()=>{const s=new Error("not implemented");return s.code="ENOSYS",s};if(!globalThis.fs){let s="";globalThis.fs={constants:{O_WRONLY:-1,O_RDWR:-1,O_CREAT:-1,O_TRUNC:-1,O_APPEND:-1,O_EXCL:-1},writeSync(e,r){s+=u.decode(r);const o=s.lastIndexOf(`
`);return o!=-1&&(console.log(s.substr(0,o)),s=s.substr(o+1)),r.length},write(e,r,o,a,d,w){if(o!==0||a!==r.length||d!==null){w(n());return}const m=this.writeSync(e,r);w(null,m)},chmod(e,r,o){o(n())},chown(e,r,o,a){a(n())},close(e,r){r(n())},fchmod(e,r,o){o(n())},fchown(e,r,o,a){a(n())},fstat(e,r){r(n())},fsync(e,r){r(null)},ftruncate(e,r,o){o(n())},lchown(e,r,o,a){a(n())},link(e,r,o){o(n())},lstat(e,r){r(n())},mkdir(e,r,o){o(n())},open(e,r,o,a){a(n())},read(e,r,o,a,d,w){w(n())},readdir(e,r){r(n())},readlink(e,r){r(n())},rename(e,r,o){o(n())},rmdir(e,r){r(n())},stat(e,r){r(n())},symlink(e,r,o){o(n())},truncate(e,r,o){o(n())},unlink(e,r){r(n())},utimes(e,r,o,a){a(n())}}}if(globalThis.process||(globalThis.process={getuid(){return-1},getgid(){return-1},geteuid(){return-1},getegid(){return-1},getgroups(){throw n()},pid:-1,ppid:-1,umask(){throw n()},cwd(){throw n()},chdir(){throw n()}}),!globalThis.crypto)throw new Error("globalThis.crypto is not available, polyfill required (crypto.getRandomValues only)");if(!globalThis.performance)throw new Error("globalThis.performance is not available, polyfill required (performance.now only)");if(!globalThis.TextEncoder)throw new Error("globalThis.TextEncoder is not available, polyfill required");if(!globalThis.TextDecoder)throw new Error("globalThis.TextDecoder is not available, polyfill required");const l=new TextEncoder("utf-8"),u=new TextDecoder("utf-8");globalThis.Go=class{constructor(){this.argv=["js"],this.env={},this.exit=t=>{t!==0&&console.warn("exit code:",t)},this._exitPromise=new Promise(t=>{this._resolveExitPromise=t}),this._pendingEvent=null,this._scheduledTimeouts=new Map,this._nextCallbackTimeoutID=1;const s=(t,i)=>{this.mem.setUint32(t+0,i,!0),this.mem.setUint32(t+4,Math.floor(i/4294967296),!0)},e=t=>{const i=this.mem.getUint32(t+0,!0),c=this.mem.getInt32(t+4,!0);return i+c*4294967296},r=t=>{const i=this.mem.getFloat64(t,!0);if(i===0)return;if(!isNaN(i))return i;const c=this.mem.getUint32(t,!0);return this._values[c]},o=(t,i)=>{if(typeof i=="number"&&i!==0){if(isNaN(i)){this.mem.setUint32(t+4,2146959360,!0),this.mem.setUint32(t,0,!0);return}this.mem.setFloat64(t,i,!0);return}if(i===void 0){this.mem.setFloat64(t,0,!0);return}let h=this._ids.get(i);h===void 0&&(h=this._idPool.pop(),h===void 0&&(h=this._values.length),this._values[h]=i,this._goRefCounts[h]=0,this._ids.set(i,h)),this._goRefCounts[h]++;let f=0;switch(typeof i){case"object":i!==null&&(f=1);break;case"string":f=2;break;case"symbol":f=3;break;case"function":f=4;break}this.mem.setUint32(t+4,2146959360|f,!0),this.mem.setUint32(t,h,!0)},a=t=>{const i=e(t+0),c=e(t+8);return new Uint8Array(this._inst.exports.mem.buffer,i,c)},d=t=>{const i=e(t+0),c=e(t+8),h=new Array(c);for(let f=0;f<c;f++)h[f]=r(i+f*8);return h},w=t=>{const i=e(t+0),c=e(t+8);return u.decode(new DataView(this._inst.exports.mem.buffer,i,c))},m=Date.now()-performance.now();this.importObject={go:{"runtime.wasmExit":t=>{t>>>=0;const i=this.mem.getInt32(t+8,!0);this.exited=!0,delete this._inst,delete this._values,delete this._goRefCounts,delete this._ids,delete this._idPool,this.exit(i)},"runtime.wasmWrite":t=>{t>>>=0;const i=e(t+8),c=e(t+16),h=this.mem.getInt32(t+24,!0);fs.writeSync(i,new Uint8Array(this._inst.exports.mem.buffer,c,h))},"runtime.resetMemoryDataView":t=>{this.mem=new DataView(this._inst.exports.mem.buffer)},"runtime.nanotime1":t=>{t>>>=0,s(t+8,(m+performance.now())*1e6)},"runtime.walltime":t=>{t>>>=0;const i=new Date().getTime();s(t+8,i/1e3),this.mem.setInt32(t+16,i%1e3*1e6,!0)},"runtime.scheduleTimeoutEvent":t=>{t>>>=0;const i=this._nextCallbackTimeoutID;this._nextCallbackTimeoutID++,this._scheduledTimeouts.set(i,setTimeout(()=>{for(this._resume();this._scheduledTimeouts.has(i);)console.warn("scheduleTimeoutEvent: missed timeout event"),this._resume()},e(t+8)+1)),this.mem.setInt32(t+16,i,!0)},"runtime.clearTimeoutEvent":t=>{t>>>=0;const i=this.mem.getInt32(t+8,!0);clearTimeout(this._scheduledTimeouts.get(i)),this._scheduledTimeouts.delete(i)},"runtime.getRandomData":t=>{t>>>=0,crypto.getRandomValues(a(t+8))},"syscall/js.finalizeRef":t=>{t>>>=0;const i=this.mem.getUint32(t+8,!0);if(this._goRefCounts[i]--,this._goRefCounts[i]===0){const c=this._values[i];this._values[i]=null,this._ids.delete(c),this._idPool.push(i)}},"syscall/js.stringVal":t=>{t>>>=0,o(t+24,w(t+8))},"syscall/js.valueGet":t=>{t>>>=0;const i=Reflect.get(r(t+8),w(t+16));t=this._inst.exports.getsp()>>>0,o(t+32,i)},"syscall/js.valueSet":t=>{t>>>=0,Reflect.set(r(t+8),w(t+16),r(t+32))},"syscall/js.valueDelete":t=>{t>>>=0,Reflect.deleteProperty(r(t+8),w(t+16))},"syscall/js.valueIndex":t=>{t>>>=0,o(t+24,Reflect.get(r(t+8),e(t+16)))},"syscall/js.valueSetIndex":t=>{t>>>=0,Reflect.set(r(t+8),e(t+16),r(t+24))},"syscall/js.valueCall":t=>{t>>>=0;try{const i=r(t+8),c=Reflect.get(i,w(t+16)),h=d(t+32),f=Reflect.apply(c,i,h);t=this._inst.exports.getsp()>>>0,o(t+56,f),this.mem.setUint8(t+64,1)}catch(i){t=this._inst.exports.getsp()>>>0,o(t+56,i),this.mem.setUint8(t+64,0)}},"syscall/js.valueInvoke":t=>{t>>>=0;try{const i=r(t+8),c=d(t+16),h=Reflect.apply(i,void 0,c);t=this._inst.exports.getsp()>>>0,o(t+40,h),this.mem.setUint8(t+48,1)}catch(i){t=this._inst.exports.getsp()>>>0,o(t+40,i),this.mem.setUint8(t+48,0)}},"syscall/js.valueNew":t=>{t>>>=0;try{const i=r(t+8),c=d(t+16),h=Reflect.construct(i,c);t=this._inst.exports.getsp()>>>0,o(t+40,h),this.mem.setUint8(t+48,1)}catch(i){t=this._inst.exports.getsp()>>>0,o(t+40,i),this.mem.setUint8(t+48,0)}},"syscall/js.valueLength":t=>{t>>>=0,s(t+16,parseInt(r(t+8).length))},"syscall/js.valuePrepareString":t=>{t>>>=0;const i=l.encode(String(r(t+8)));o(t+16,i),s(t+24,i.length)},"syscall/js.valueLoadString":t=>{t>>>=0;const i=r(t+8);a(t+16).set(i)},"syscall/js.valueInstanceOf":t=>{t>>>=0,this.mem.setUint8(t+24,r(t+8)instanceof r(t+16)?1:0)},"syscall/js.copyBytesToGo":t=>{t>>>=0;const i=a(t+8),c=r(t+32);if(!(c instanceof Uint8Array||c instanceof Uint8ClampedArray)){this.mem.setUint8(t+48,0);return}const h=c.subarray(0,i.length);i.set(h),s(t+40,h.length),this.mem.setUint8(t+48,1)},"syscall/js.copyBytesToJS":t=>{t>>>=0;const i=r(t+8),c=a(t+16);if(!(i instanceof Uint8Array||i instanceof Uint8ClampedArray)){this.mem.setUint8(t+48,0);return}const h=c.subarray(0,i.length);i.set(h),s(t+40,h.length),this.mem.setUint8(t+48,1)},debug:t=>{console.log(t)}}}}async run(s){if(!(s instanceof WebAssembly.Instance))throw new Error("Go.run: WebAssembly.Instance expected");this._inst=s,this.mem=new DataView(this._inst.exports.mem.buffer),this._values=[NaN,0,null,!0,!1,globalThis,this],this._goRefCounts=new Array(this._values.length).fill(1/0),this._ids=new Map([[0,1],[null,2],[!0,3],[!1,4],[globalThis,5],[this,6]]),this._idPool=[],this.exited=!1;let e=4096;const r=t=>{const i=e,c=l.encode(t+"\0");return new Uint8Array(this.mem.buffer,e,c.length).set(c),e+=c.length,e%8!==0&&(e+=8-e%8),i},o=this.argv.length,a=[];this.argv.forEach(t=>{a.push(r(t))}),a.push(0),Object.keys(this.env).sort().forEach(t=>{a.push(r(`${t}=${this.env[t]}`))}),a.push(0);const w=e;a.forEach(t=>{this.mem.setUint32(e,t,!0),this.mem.setUint32(e+4,0,!0),e+=8});const m=4096+8192;if(e>=m)throw new Error("total length of command line and environment variables exceeds limit");this._inst.exports.run(o,w),this.exited&&this._resolveExitPromise(),await this._exitPromise}_resume(){if(this.exited)throw new Error("Go program has already exited");this._inst.exports.resume(),this.exited&&this._resolveExitPromise()}_makeFuncWrapper(s){const e=this;return function(){const r={id:s,this:this,args:arguments};return e._pendingEvent=r,e._resume(),r.result}}}})();function H(n){let l;const u=n[2].default,s=N(u,n,n[1],null);return{c(){s&&s.c()},l(e){s&&s.l(e)},m(e,r){s&&s.m(e,r),l=!0},p(e,r){s&&s.p&&(!l||r&2)&&P(s,u,e,e[1],l?V(u,e[1],r,null):O(e[1]),null)},i(e){l||(T(s,e),l=!0)},o(e){k(s,e),l=!1},d(e){s&&s.d(e)}}}function L(n){let l,u,s,e,r,o,a,d,w;return{c(){l=b("div"),u=b("div"),s=b("div"),e=b("div"),r=b("h1"),o=j("Timeless Calculator"),a=F(),d=b("h2"),w=j("Loading..."),this.h()},l(m){l=y(m,"DIV",{class:!0});var t=x(l);u=y(t,"DIV",{class:!0});var i=x(u);s=y(i,"DIV",{class:!0});var c=x(s);e=y(c,"DIV",{});var h=x(e);r=y(h,"H1",{class:!0});var f=x(r);o=I(f,"Timeless Calculator"),f.forEach(g),a=M(h),d=y(h,"H2",{class:!0});var E=x(d);w=I(E,"Loading..."),E.forEach(g),h.forEach(g),c.forEach(g),i.forEach(g),t.forEach(g),this.h()},h(){v(r,"class","text-white mb-10 text-center"),v(d,"class","text-center"),v(s,"class","py-10 flex flex-col justify-between"),v(u,"class","flex flex-col"),v(l,"class","flex flex-row justify-center h-screen")},m(m,t){z(m,l,t),p(l,u),p(u,s),p(s,e),p(e,r),p(r,o),p(e,a),p(e,d),p(d,w)},p:_,i:_,o:_,d(m){m&&g(l)}}}function $(n){let l,u,s,e;const r=[L,H],o=[];function a(d,w){return d[0]?0:1}return l=a(n),u=o[l]=r[l](n),{c(){u.c(),s=U()},l(d){u.l(d),s=U()},m(d,w){o[l].m(d,w),z(d,s,w),e=!0},p(d,[w]){let m=l;l=a(d),l===m?o[l].p(d,w):(A(),k(o[m],1,1,()=>{o[m]=null}),D(),u=o[l],u?u.p(d,w):(u=o[l]=r[l](d),u.c()),T(u,1),u.m(s.parentNode,s))},i(d){e||(T(u),e=!0)},o(d){k(u),e=!1},d(d){o[l].d(d),d&&g(s)}}}function q(n,l,u){let{$$slots:s={},$$scope:e}=l,r=!0;const o=new Go;return WebAssembly.instantiateStreaming(fetch(G+"/calculator.wasm"),o.importObject).then(a=>{o.run(a.instance),u(0,r=!1)}),n.$$set=a=>{"$$scope"in a&&u(1,e=a.$$scope)},[r,e,s]}class J extends R{constructor(l){super(),C(this,l,q,$,S,{})}}export{J as default};
