function k(){}function tt(t,e){for(const n in e)t[n]=e[n];return t}function J(t){return t()}function W(){return Object.create(null)}function v(t){t.forEach(J)}function z(t){return typeof t=="function"}function wt(t,e){return t!=t?e==e:t!==e||t&&typeof t=="object"||typeof t=="function"}function et(t){return Object.keys(t).length===0}function nt(t,...e){if(t==null)return k;const n=t.subscribe(...e);return n.unsubscribe?()=>n.unsubscribe():n}function vt(t,e,n){t.$$.on_destroy.push(nt(e,n))}function Et(t,e,n,i){if(t){const r=K(t,e,n,i);return t[0](r)}}function K(t,e,n,i){return t[1]&&i?tt(n.ctx.slice(),t[1](i(e))):n.ctx}function kt(t,e,n,i){if(t[2]&&i){const r=t[2](i(n));if(e.dirty===void 0)return r;if(typeof r=="object"){const u=[],s=Math.max(e.dirty.length,r.length);for(let o=0;o<s;o+=1)u[o]=e.dirty[o]|r[o];return u}return e.dirty|r}return e.dirty}function Nt(t,e,n,i,r,u){if(r){const s=K(e,n,i,u);t.p(s,r)}}function St(t){if(t.ctx.length>32){const e=[],n=t.ctx.length/32;for(let i=0;i<n;i++)e[i]=-1;return e}return-1}function At(t){const e={};for(const n in t)e[n]=!0;return e}function Ct(t){return t&&z(t.destroy)?t.destroy:k}let j=!1;function it(){j=!0}function rt(){j=!1}function st(t,e,n,i){for(;t<e;){const r=t+(e-t>>1);n(r)<=i?t=r+1:e=r}return t}function ct(t){if(t.hydrate_init)return;t.hydrate_init=!0;let e=t.childNodes;if(t.nodeName==="HEAD"){const c=[];for(let l=0;l<e.length;l++){const f=e[l];f.claim_order!==void 0&&c.push(f)}e=c}const n=new Int32Array(e.length+1),i=new Int32Array(e.length);n[0]=-1;let r=0;for(let c=0;c<e.length;c++){const l=e[c].claim_order,f=(r>0&&e[n[r]].claim_order<=l?r+1:st(1,r,d=>e[n[d]].claim_order,l))-1;i[c]=n[f]+1;const _=f+1;n[_]=c,r=Math.max(_,r)}const u=[],s=[];let o=e.length-1;for(let c=n[r]+1;c!=0;c=i[c-1]){for(u.push(e[c-1]);o>=c;o--)s.push(e[o]);o--}for(;o>=0;o--)s.push(e[o]);u.reverse(),s.sort((c,l)=>c.claim_order-l.claim_order);for(let c=0,l=0;c<s.length;c++){for(;l<u.length&&s[c].claim_order>=u[l].claim_order;)l++;const f=l<u.length?u[l]:null;t.insertBefore(s[c],f)}}function ot(t,e){if(j){for(ct(t),(t.actual_end_child===void 0||t.actual_end_child!==null&&t.actual_end_child.parentNode!==t)&&(t.actual_end_child=t.firstChild);t.actual_end_child!==null&&t.actual_end_child.claim_order===void 0;)t.actual_end_child=t.actual_end_child.nextSibling;e!==t.actual_end_child?(e.claim_order!==void 0||e.parentNode!==t)&&t.insertBefore(e,t.actual_end_child):t.actual_end_child=e.nextSibling}else(e.parentNode!==t||e.nextSibling!==null)&&t.appendChild(e)}function jt(t,e,n){j&&!n?ot(t,e):(e.parentNode!==t||e.nextSibling!=n)&&t.insertBefore(e,n||null)}function lt(t){t.parentNode&&t.parentNode.removeChild(t)}function Mt(t,e){for(let n=0;n<t.length;n+=1)t[n]&&t[n].d(e)}function ut(t){return document.createElement(t)}function at(t){return document.createElementNS("http://www.w3.org/2000/svg",t)}function F(t){return document.createTextNode(t)}function Ot(){return F(" ")}function Pt(){return F("")}function Tt(t,e,n,i){return t.addEventListener(e,n,i),()=>t.removeEventListener(e,n,i)}function Dt(t){return function(e){return e.preventDefault(),t.call(this,e)}}function Lt(t){return function(e){return e.stopPropagation(),t.call(this,e)}}function ft(t,e,n){n==null?t.removeAttribute(e):t.getAttribute(e)!==n&&t.setAttribute(e,n)}function Bt(t,e){const n=Object.getOwnPropertyDescriptors(t.__proto__);for(const i in e)e[i]==null?t.removeAttribute(i):i==="style"?t.style.cssText=e[i]:i==="__value"?t.value=t[i]=e[i]:n[i]&&n[i].set?t[i]=e[i]:ft(t,i,e[i])}function qt(t){return t===""?null:+t}function dt(t){return Array.from(t.childNodes)}function _t(t){t.claim_info===void 0&&(t.claim_info={last_index:0,total_claimed:0})}function Q(t,e,n,i,r=!1){_t(t);const u=(()=>{for(let s=t.claim_info.last_index;s<t.length;s++){const o=t[s];if(e(o)){const c=n(o);return c===void 0?t.splice(s,1):t[s]=c,r||(t.claim_info.last_index=s),o}}for(let s=t.claim_info.last_index-1;s>=0;s--){const o=t[s];if(e(o)){const c=n(o);return c===void 0?t.splice(s,1):t[s]=c,r?c===void 0&&t.claim_info.last_index--:t.claim_info.last_index=s,o}}return i()})();return u.claim_order=t.claim_info.total_claimed,t.claim_info.total_claimed+=1,u}function R(t,e,n,i){return Q(t,r=>r.nodeName===e,r=>{const u=[];for(let s=0;s<r.attributes.length;s++){const o=r.attributes[s];n[o.name]||u.push(o.name)}u.forEach(s=>r.removeAttribute(s))},()=>i(e))}function zt(t,e,n){return R(t,e,n,ut)}function Ft(t,e,n){return R(t,e,n,at)}function ht(t,e){return Q(t,n=>n.nodeType===3,n=>{const i=""+e;if(n.data.startsWith(i)){if(n.data.length!==i.length)return n.splitText(i.length)}else n.data=i},()=>F(e),!0)}function Ht(t){return ht(t," ")}function It(t,e){e=""+e,t.data!==e&&(t.data=e)}function Ut(t,e){t.value=e??""}function Wt(t,e,n,i){n===null?t.style.removeProperty(e):t.style.setProperty(e,n,i?"important":"")}function Gt(t,e,n){t.classList[n?"add":"remove"](e)}function mt(t,e,{bubbles:n=!1,cancelable:i=!1}={}){const r=document.createEvent("CustomEvent");return r.initCustomEvent(t,n,i,e),r}function Jt(t,e){return new t(e)}let N;function E(t){N=t}function b(){if(!N)throw new Error("Function called outside component initialization");return N}function Kt(t){b().$$.before_update.push(t)}function Qt(t){b().$$.on_mount.push(t)}function Rt(t){b().$$.after_update.push(t)}function Vt(t){b().$$.on_destroy.push(t)}function Xt(){const t=b();return(e,n,{cancelable:i=!1}={})=>{const r=t.$$.callbacks[e];if(r){const u=mt(e,n,{cancelable:i});return r.slice().forEach(s=>{s.call(t,u)}),!u.defaultPrevented}return!0}}function Yt(t,e){return b().$$.context.set(t,e),e}function Zt(t){return b().$$.context.get(t)}function te(t,e){const n=t.$$.callbacks[e.type];n&&n.slice().forEach(i=>i.call(this,e))}const x=[],G=[];let w=[];const L=[],V=Promise.resolve();let B=!1;function X(){B||(B=!0,V.then(Y))}function ee(){return X(),V}function q(t){w.push(t)}function ne(t){L.push(t)}const D=new Set;let $=0;function Y(){if($!==0)return;const t=N;do{try{for(;$<x.length;){const e=x[$];$++,E(e),pt(e.$$)}}catch(e){throw x.length=0,$=0,e}for(E(null),x.length=0,$=0;G.length;)G.pop()();for(let e=0;e<w.length;e+=1){const n=w[e];D.has(n)||(D.add(n),n())}w.length=0}while(x.length);for(;L.length;)L.pop()();B=!1,D.clear(),E(t)}function pt(t){if(t.fragment!==null){t.update(),v(t.before_update);const e=t.dirty;t.dirty=[-1],t.fragment&&t.fragment.p(t.ctx,e),t.after_update.forEach(q)}}function yt(t){const e=[],n=[];w.forEach(i=>t.indexOf(i)===-1?e.push(i):n.push(i)),n.forEach(i=>i()),w=e}const C=new Set;let g;function ie(){g={r:0,c:[],p:g}}function re(){g.r||v(g.c),g=g.p}function Z(t,e){t&&t.i&&(C.delete(t),t.i(e))}function gt(t,e,n,i){if(t&&t.o){if(C.has(t))return;C.add(t),g.c.push(()=>{C.delete(t),i&&(n&&t.d(1),i())}),t.o(e)}else i&&i()}const se=typeof window<"u"?window:typeof globalThis<"u"?globalThis:global;function ce(t,e){gt(t,1,1,()=>{e.delete(t.key)})}function oe(t,e,n,i,r,u,s,o,c,l,f,_){let d=t.length,m=u.length,h=d;const M={};for(;h--;)M[t[h].key]=h;const S=[],O=new Map,P=new Map,H=[];for(h=m;h--;){const a=_(r,u,h),p=n(a);let y=s.get(p);y?i&&H.push(()=>y.p(a,e)):(y=l(p,a),y.c()),O.set(p,S[h]=y),p in M&&P.set(p,Math.abs(h-M[p]))}const I=new Set,U=new Set;function T(a){Z(a,1),a.m(o,f),s.set(a.key,a),f=a.first,m--}for(;d&&m;){const a=S[m-1],p=t[d-1],y=a.key,A=p.key;a===p?(f=a.first,d--,m--):O.has(A)?!s.has(y)||I.has(y)?T(a):U.has(A)?d--:P.get(y)>P.get(A)?(U.add(y),T(a)):(I.add(A),d--):(c(p,s),d--)}for(;d--;){const a=t[d];O.has(a.key)||c(a,s)}for(;m;)T(S[m-1]);return v(H),S}function le(t,e){const n={},i={},r={$$scope:1};let u=t.length;for(;u--;){const s=t[u],o=e[u];if(o){for(const c in s)c in o||(i[c]=1);for(const c in o)r[c]||(n[c]=o[c],r[c]=1);t[u]=o}else for(const c in s)r[c]=1}for(const s in i)s in n||(n[s]=void 0);return n}function ue(t,e,n){const i=t.$$.props[e];i!==void 0&&(t.$$.bound[i]=n,n(t.$$.ctx[i]))}function ae(t){t&&t.c()}function fe(t,e){t&&t.l(e)}function bt(t,e,n,i){const{fragment:r,after_update:u}=t.$$;r&&r.m(e,n),i||q(()=>{const s=t.$$.on_mount.map(J).filter(z);t.$$.on_destroy?t.$$.on_destroy.push(...s):v(s),t.$$.on_mount=[]}),u.forEach(q)}function $t(t,e){const n=t.$$;n.fragment!==null&&(yt(n.after_update),v(n.on_destroy),n.fragment&&n.fragment.d(e),n.on_destroy=n.fragment=null,n.ctx=[])}function xt(t,e){t.$$.dirty[0]===-1&&(x.push(t),X(),t.$$.dirty.fill(0)),t.$$.dirty[e/31|0]|=1<<e%31}function de(t,e,n,i,r,u,s,o=[-1]){const c=N;E(t);const l=t.$$={fragment:null,ctx:[],props:u,update:k,not_equal:r,bound:W(),on_mount:[],on_destroy:[],on_disconnect:[],before_update:[],after_update:[],context:new Map(e.context||(c?c.$$.context:[])),callbacks:W(),dirty:o,skip_bound:!1,root:e.target||c.$$.root};s&&s(l.root);let f=!1;if(l.ctx=n?n(t,e.props||{},(_,d,...m)=>{const h=m.length?m[0]:d;return l.ctx&&r(l.ctx[_],l.ctx[_]=h)&&(!l.skip_bound&&l.bound[_]&&l.bound[_](h),f&&xt(t,_)),d}):[],l.update(),f=!0,v(l.before_update),l.fragment=i?i(l.ctx):!1,e.target){if(e.hydrate){it();const _=dt(e.target);l.fragment&&l.fragment.l(_),_.forEach(lt)}else l.fragment&&l.fragment.c();e.intro&&Z(t.$$.fragment),bt(t,e.target,e.anchor,e.customElement),rt(),Y()}E(c)}class _e{$destroy(){$t(this,1),this.$destroy=k}$on(e,n){if(!z(n))return k;const i=this.$$.callbacks[e]||(this.$$.callbacks[e]=[]);return i.push(n),()=>{const r=i.indexOf(n);r!==-1&&i.splice(r,1)}}$set(e){this.$$set&&!et(e)&&(this.$$.skip_bound=!0,this.$$set(e),this.$$.skip_bound=!1)}}export{Gt as $,bt as A,$t as B,Et as C,Nt as D,St as E,kt as F,ot as G,k as H,vt as I,ue as J,ne as K,Ut as L,Tt as M,qt as N,v as O,Mt as P,Dt as Q,z as R,_e as S,Yt as T,Vt as U,Zt as V,te as W,Xt as X,se as Y,oe as Z,ce as _,Ot as a,at as a0,Ft as a1,tt as a2,Bt as a3,Ct as a4,le as a5,At as a6,Kt as a7,Lt as a8,nt as a9,jt as b,Ht as c,gt as d,Pt as e,re as f,Z as g,lt as h,de as i,Rt as j,ut as k,zt as l,dt as m,ft as n,Qt as o,Wt as p,F as q,ht as r,wt as s,ee as t,It as u,ie as v,G as w,Jt as x,ae as y,fe as z};