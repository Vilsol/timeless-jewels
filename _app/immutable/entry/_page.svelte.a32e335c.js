import{S as Te,i as je,s as Je,w as K,J as z,k as I,q as S,a as $,y as Q,l as P,m as E,r as g,h as p,c as D,z as F,n as A,b as T,G as d,A as W,K as X,g as C,d as G,f as Ie,B as Y,I as $e,e as N,v as Pe,L as oe,M as fe,N as Ee,O as De,u as q,P as Z}from"../chunks/index.62a09d90.js";import{S as x,g as Oe}from"../chunks/navigation.26e81368.js";import{p as Ve}from"../chunks/stores.a1ce8246.js";import{b as qe}from"../chunks/paths.732429ef.js";import{d as J,c as Me}from"../chunks/index.f119c485.js";function re(n,e,a){const l=n.slice();return l[15]=e[a],l}function ue(n,e,a){const l=n.slice();l[18]=e[a],l[21]=a;const i=J.GetStatByIndex(l[15].AlternatePassiveAddition.StatsKeys[l[21]]);return l[19]=i,l}function ce(n,e,a){const l=n.slice();l[18]=e[a],l[21]=a;const i=J.GetStatByIndex(l[4].AlternatePassiveSkill.StatsKeys[l[21]]);return l[19]=i,l}function _e(n){let e,a,l,i,s,t,o,r=n[1]&&Object.keys(J.TimelessJewelConquerors[n[0].value]).indexOf(n[1].value)>=0,v,c;function h(u){n[10](u)}let m={items:n[5]};n[1]!==void 0&&(m.value=n[1]),s=new x({props:m}),K.push(()=>z(s,"value",h)),s.$on("select",n[8]);let f=r&&he(n);return{c(){e=I("div"),a=I("h3"),l=S("Conqueror"),i=$(),Q(s.$$.fragment),o=$(),f&&f.c(),v=N(),this.h()},l(u){e=P(u,"DIV",{class:!0});var _=E(e);a=P(_,"H3",{class:!0});var k=E(a);l=g(k,"Conqueror"),k.forEach(p),i=D(_),F(s.$$.fragment,_),_.forEach(p),o=D(u),f&&f.l(u),v=N(),this.h()},h(){A(a,"class","mb-2"),A(e,"class","mt-4")},m(u,_){T(u,e,_),d(e,a),d(a,l),d(e,i),W(s,e,null),T(u,o,_),f&&f.m(u,_),T(u,v,_),c=!0},p(u,_){const k={};_&32&&(k.items=u[5]),!t&&_&2&&(t=!0,k.value=u[1],X(()=>t=!1)),s.$set(k),_&3&&(r=u[1]&&Object.keys(J.TimelessJewelConquerors[u[0].value]).indexOf(u[1].value)>=0),r?f?(f.p(u,_),_&3&&C(f,1)):(f=he(u),f.c(),C(f,1),f.m(v.parentNode,v)):f&&(Pe(),G(f,1,1,()=>{f=null}),Ie())},i(u){c||(C(s.$$.fragment,u),C(f),c=!0)},o(u){G(s.$$.fragment,u),G(f),c=!1},d(u){u&&p(e),Y(s),u&&p(o),f&&f.d(u),u&&p(v)}}}function he(n){let e,a,l,i,s,t,o,r,v;function c(f){n[11](f)}let h={items:n[7]};n[2]!==void 0&&(h.value=n[2]),s=new x({props:h}),K.push(()=>z(s,"value",c)),s.$on("select",n[8]);let m=n[2]&&ve(n);return{c(){e=I("div"),a=I("h3"),l=S("Passive Skill"),i=$(),Q(s.$$.fragment),o=$(),m&&m.c(),r=N(),this.h()},l(f){e=P(f,"DIV",{class:!0});var u=E(e);a=P(u,"H3",{class:!0});var _=E(a);l=g(_,"Passive Skill"),_.forEach(p),i=D(u),F(s.$$.fragment,u),u.forEach(p),o=D(f),m&&m.l(f),r=N(),this.h()},h(){A(a,"class","mb-2"),A(e,"class","mt-4")},m(f,u){T(f,e,u),d(e,a),d(a,l),d(e,i),W(s,e,null),T(f,o,u),m&&m.m(f,u),T(f,r,u),v=!0},p(f,u){const _={};!t&&u&4&&(t=!0,_.value=f[2],X(()=>t=!1)),s.$set(_),f[2]?m?m.p(f,u):(m=ve(f),m.c(),m.m(r.parentNode,r)):m&&(m.d(1),m=null)},i(f){v||(C(s.$$.fragment,f),v=!0)},o(f){G(s.$$.fragment,f),v=!1},d(f){f&&p(e),Y(s),f&&p(o),m&&m.d(f),f&&p(r)}}}function ve(n){let e,a,l,i,s,t,o,r,v,c,h,m,f=(n[3]<J.TimelessJewelSeedRanges[n[0].value].Min||n[3]>J.TimelessJewelSeedRanges[n[0].value].Max)&&de(n),u=n[4]&&me(n);return{c(){e=I("div"),a=I("h3"),l=S("Seed"),i=$(),s=I("input"),r=$(),f&&f.c(),v=$(),u&&u.c(),c=N(),this.h()},l(_){e=P(_,"DIV",{class:!0});var k=E(e);a=P(k,"H3",{class:!0});var V=E(a);l=g(V,"Seed"),V.forEach(p),i=D(k),s=P(k,"INPUT",{type:!0,class:!0,min:!0,max:!0}),r=D(k),f&&f.l(k),k.forEach(p),v=D(_),u&&u.l(_),c=N(),this.h()},h(){A(a,"class","mb-2"),A(s,"type","number"),A(s,"class","seed"),A(s,"min",t=J.TimelessJewelSeedRanges[n[0].value].Min),A(s,"max",o=J.TimelessJewelSeedRanges[n[0].value].Max),A(e,"class","mt-4")},m(_,k){T(_,e,k),d(e,a),d(a,l),d(e,i),d(e,s),oe(s,n[3]),d(e,r),f&&f.m(e,null),T(_,v,k),u&&u.m(_,k),T(_,c,k),h||(m=[fe(s,"input",n[12]),fe(s,"blur",n[8])],h=!0)},p(_,k){k&1&&t!==(t=J.TimelessJewelSeedRanges[_[0].value].Min)&&A(s,"min",t),k&1&&o!==(o=J.TimelessJewelSeedRanges[_[0].value].Max)&&A(s,"max",o),k&8&&Ee(s.value)!==_[3]&&oe(s,_[3]),_[3]<J.TimelessJewelSeedRanges[_[0].value].Min||_[3]>J.TimelessJewelSeedRanges[_[0].value].Max?f?f.p(_,k):(f=de(_),f.c(),f.m(e,null)):f&&(f.d(1),f=null),_[4]?u?u.p(_,k):(u=me(_),u.c(),u.m(c.parentNode,c)):u&&(u.d(1),u=null)},d(_){_&&p(e),f&&f.d(),_&&p(v),u&&u.d(_),_&&p(c),h=!1,De(m)}}}function de(n){let e,a,l=J.TimelessJewelSeedRanges[n[0].value].Min+"",i,s,t=J.TimelessJewelSeedRanges[n[0].value].Max+"",o;return{c(){e=I("div"),a=S("Seed must be between "),i=S(l),s=S(" and "),o=S(t),this.h()},l(r){e=P(r,"DIV",{class:!0});var v=E(e);a=g(v,"Seed must be between "),i=g(v,l),s=g(v," and "),o=g(v,t),v.forEach(p),this.h()},h(){A(e,"class","mt-2")},m(r,v){T(r,e,v),d(e,a),d(e,i),d(e,s),d(e,o)},p(r,v){v&1&&l!==(l=J.TimelessJewelSeedRanges[r[0].value].Min+"")&&q(i,l),v&1&&t!==(t=J.TimelessJewelSeedRanges[r[0].value].Max+"")&&q(o,t)},d(r){r&&p(e)}}}function me(n){var s;let e,a,l=n[4].AlternatePassiveSkill&&pe(n),i="AlternatePassiveAdditionInformations"in n[4]&&((s=n[4].AlternatePassiveAdditionInformations)==null?void 0:s.length)>0&&Se(n);return{c(){l&&l.c(),e=$(),i&&i.c(),a=N()},l(t){l&&l.l(t),e=D(t),i&&i.l(t),a=N()},m(t,o){l&&l.m(t,o),T(t,e,o),i&&i.m(t,o),T(t,a,o)},p(t,o){var r;t[4].AlternatePassiveSkill?l?l.p(t,o):(l=pe(t),l.c(),l.m(e.parentNode,e)):l&&(l.d(1),l=null),"AlternatePassiveAdditionInformations"in t[4]&&((r=t[4].AlternatePassiveAdditionInformations)==null?void 0:r.length)>0?i?i.p(t,o):(i=Se(t),i.c(),i.m(a.parentNode,a)):i&&(i.d(1),i=null)},d(t){l&&l.d(t),t&&p(e),i&&i.d(t),t&&p(a)}}}function pe(n){let e,a,l,i,s,t=n[4].AlternatePassiveSkill.Name+"",o,r,v=n[4].AlternatePassiveSkill.ID+"",c,h,m=n[4].AlternatePassiveSkill+"",f,u,_,k=n[4].StatRolls&&Object.keys(n[4].StatRolls).length>0,V,b=k&&be(n);return{c(){e=I("div"),a=I("h3"),l=S("Alternate Passive Skill"),i=$(),s=I("span"),o=S(t),r=S(" ("),c=S(v),h=S(") ("),f=S(m),u=S(")"),_=$(),b&&b.c(),V=N(),this.h()},l(w){e=P(w,"DIV",{class:!0});var R=E(e);a=P(R,"H3",{});var y=E(a);l=g(y,"Alternate Passive Skill"),y.forEach(p),i=D(R),s=P(R,"SPAN",{});var M=E(s);o=g(M,t),r=g(M," ("),c=g(M,v),h=g(M,") ("),f=g(M,m),u=g(M,")"),M.forEach(p),R.forEach(p),_=D(w),b&&b.l(w),V=N(),this.h()},h(){A(e,"class","mt-4")},m(w,R){T(w,e,R),d(e,a),d(a,l),d(e,i),d(e,s),d(s,o),d(s,r),d(s,c),d(s,h),d(s,f),d(s,u),T(w,_,R),b&&b.m(w,R),T(w,V,R)},p(w,R){R&16&&t!==(t=w[4].AlternatePassiveSkill.Name+"")&&q(o,t),R&16&&v!==(v=w[4].AlternatePassiveSkill.ID+"")&&q(c,v),R&16&&m!==(m=w[4].AlternatePassiveSkill+"")&&q(f,m),R&16&&(k=w[4].StatRolls&&Object.keys(w[4].StatRolls).length>0),k?b?b.p(w,R):(b=be(w),b.c(),b.m(V.parentNode,V)):b&&(b.d(1),b=null)},d(w){w&&p(e),w&&p(_),b&&b.d(w),w&&p(V)}}}function be(n){let e,a=Object.keys(n[4].StatRolls),l=[];for(let i=0;i<a.length;i+=1)l[i]=ke(ce(n,a,i));return{c(){e=I("ol");for(let i=0;i<l.length;i+=1)l[i].c();this.h()},l(i){e=P(i,"OL",{class:!0});var s=E(e);for(let t=0;t<l.length;t+=1)l[t].l(s);s.forEach(p),this.h()},h(){A(e,"class","mt-4 list-decimal pl-8")},m(i,s){T(i,e,s);for(let t=0;t<l.length;t+=1)l[t]&&l[t].m(e,null)},p(i,s){if(s&16){a=Object.keys(i[4].StatRolls);let t;for(t=0;t<a.length;t+=1){const o=ce(i,a,t);l[t]?l[t].p(o,s):(l[t]=ke(o),l[t].c(),l[t].m(e,null))}for(;t<l.length;t+=1)l[t].d(1);l.length=a.length}},d(i){i&&p(e),Z(l,i)}}}function ke(n){let e,a=(n[19].Text||"<no name>")+"",l,i,s=n[19].ID+"",t,o,r=n[4].StatRolls[n[18]]+"",v;return{c(){e=I("li"),l=S(a),i=S(" ("),t=S(s),o=S(") - "),v=S(r)},l(c){e=P(c,"LI",{});var h=E(e);l=g(h,a),i=g(h," ("),t=g(h,s),o=g(h,") - "),v=g(h,r),h.forEach(p)},m(c,h){T(c,e,h),d(e,l),d(e,i),d(e,t),d(e,o),d(e,v)},p(c,h){h&16&&a!==(a=(c[19].Text||"<no name>")+"")&&q(l,a),h&16&&s!==(s=c[19].ID+"")&&q(t,s),h&16&&r!==(r=c[4].StatRolls[c[18]]+"")&&q(v,r)},d(c){c&&p(e)}}}function Se(n){let e,a,l,i,s,t=n[4].AlternatePassiveAdditionInformations,o=[];for(let r=0;r<t.length;r+=1)o[r]=Ae(re(n,t,r));return{c(){e=I("div"),a=I("h3"),l=S("Additions"),i=$(),s=I("ul");for(let r=0;r<o.length;r+=1)o[r].c();this.h()},l(r){e=P(r,"DIV",{class:!0});var v=E(e);a=P(v,"H3",{});var c=E(a);l=g(c,"Additions"),c.forEach(p),i=D(v),s=P(v,"UL",{class:!0});var h=E(s);for(let m=0;m<o.length;m+=1)o[m].l(h);h.forEach(p),v.forEach(p),this.h()},h(){A(s,"class","list-disc pl-8"),A(e,"class","mt-4")},m(r,v){T(r,e,v),d(e,a),d(a,l),d(e,i),d(e,s);for(let c=0;c<o.length;c+=1)o[c]&&o[c].m(s,null)},p(r,v){if(v&16){t=r[4].AlternatePassiveAdditionInformations;let c;for(c=0;c<t.length;c+=1){const h=re(r,t,c);o[c]?o[c].p(h,v):(o[c]=Ae(h),o[c].c(),o[c].m(s,null))}for(;c<o.length;c+=1)o[c].d(1);o.length=t.length}},d(r){r&&p(e),Z(o,r)}}}function ge(n){let e,a=Object.keys(n[15].StatRolls),l=[];for(let i=0;i<a.length;i+=1)l[i]=we(ue(n,a,i));return{c(){e=I("ol");for(let i=0;i<l.length;i+=1)l[i].c();this.h()},l(i){e=P(i,"OL",{class:!0});var s=E(e);for(let t=0;t<l.length;t+=1)l[t].l(s);s.forEach(p),this.h()},h(){A(e,"class","list-decimal pl-8")},m(i,s){T(i,e,s);for(let t=0;t<l.length;t+=1)l[t]&&l[t].m(e,null)},p(i,s){if(s&16){a=Object.keys(i[15].StatRolls);let t;for(t=0;t<a.length;t+=1){const o=ue(i,a,t);l[t]?l[t].p(o,s):(l[t]=we(o),l[t].c(),l[t].m(e,null))}for(;t<l.length;t+=1)l[t].d(1);l.length=a.length}},d(i){i&&p(e),Z(l,i)}}}function we(n){let e,a=(n[19].Text||"<no name>")+"",l,i,s=n[19].ID+"",t,o,r=n[15].StatRolls[n[18]]+"",v;return{c(){e=I("li"),l=S(a),i=S(" ("),t=S(s),o=S(") - "),v=S(r)},l(c){e=P(c,"LI",{});var h=E(e);l=g(h,a),i=g(h," ("),t=g(h,s),o=g(h,") - "),v=g(h,r),h.forEach(p)},m(c,h){T(c,e,h),d(e,l),d(e,i),d(e,t),d(e,o),d(e,v)},p(c,h){h&16&&a!==(a=(c[19].Text||"<no name>")+"")&&q(l,a),h&16&&s!==(s=c[19].ID+"")&&q(t,s),h&16&&r!==(r=c[15].StatRolls[c[18]]+"")&&q(v,r)},d(c){c&&p(e)}}}function Ae(n){let e,a,l=n[15].AlternatePassiveAddition.ID+"",i,s,t=n[15].AlternatePassiveAddition.Index+"",o,r,v,c=n[15].StatRolls&&Object.keys(n[15].StatRolls).length>0,h,m=c&&ge(n);return{c(){e=I("li"),a=I("span"),i=S(l),s=S(" ("),o=S(t),r=S(")"),v=$(),m&&m.c(),h=$(),this.h()},l(f){e=P(f,"LI",{class:!0});var u=E(e);a=P(u,"SPAN",{});var _=E(a);i=g(_,l),s=g(_," ("),o=g(_,t),r=g(_,")"),_.forEach(p),v=D(u),m&&m.l(u),h=D(u),u.forEach(p),this.h()},h(){A(e,"class","mt-4")},m(f,u){T(f,e,u),d(e,a),d(a,i),d(a,s),d(a,o),d(a,r),d(e,v),m&&m.m(e,null),d(e,h)},p(f,u){u&16&&l!==(l=f[15].AlternatePassiveAddition.ID+"")&&q(i,l),u&16&&t!==(t=f[15].AlternatePassiveAddition.Index+"")&&q(o,t),u&16&&(c=f[15].StatRolls&&Object.keys(f[15].StatRolls).length>0),c?m?m.p(f,u):(m=ge(f),m.c(),m.m(e,h)):m&&(m.d(1),m=null)},d(f){f&&p(e),m&&m.d()}}}function Ne(n){let e,a,l,i,s,t,o,r,v,c,h,m,f,u,_,k,V,b,w,R,y,M;function Re(O){n[9](O)}let ee={items:n[6]};n[0]!==void 0&&(ee.value=n[0]),_=new x({props:ee}),K.push(()=>z(_,"value",Re)),_.$on("select",n[8]);let j=n[0]&&_e(n);return{c(){e=I("div"),a=I("div"),l=I("div"),i=I("h1"),s=S("Timeless Calculator"),t=$(),o=I("a"),r=I("h2"),v=S("Skill Tree View"),c=$(),h=I("div"),m=I("h3"),f=S("Timeless Jewel"),u=$(),Q(_.$$.fragment),V=$(),j&&j.c(),b=$(),w=I("div"),R=I("a"),y=S("Source (Github)"),this.h()},l(O){e=P(O,"DIV",{class:!0});var H=E(e);a=P(H,"DIV",{class:!0});var L=E(a);l=P(L,"DIV",{});var U=E(l);i=P(U,"H1",{class:!0});var le=E(i);s=g(le,"Timeless Calculator"),le.forEach(p),t=D(U),o=P(U,"A",{href:!0});var te=E(o);r=P(te,"H2",{class:!0});var se=E(r);v=g(se,"Skill Tree View"),se.forEach(p),te.forEach(p),c=D(U),h=P(U,"DIV",{class:!0});var B=E(h);m=P(B,"H3",{class:!0});var ie=E(m);f=g(ie,"Timeless Jewel"),ie.forEach(p),u=D(B),F(_.$$.fragment,B),V=D(B),j&&j.l(B),B.forEach(p),U.forEach(p),b=D(L),w=P(L,"DIV",{class:!0});var ae=E(w);R=P(ae,"A",{href:!0,target:!0,rel:!0});var ne=E(R);y=g(ne,"Source (Github)"),ne.forEach(p),ae.forEach(p),L.forEach(p),H.forEach(p),this.h()},h(){A(i,"class","text-white mb-10 text-center"),A(r,"class","text-white mb-10 text-center underline text-orange-500"),A(o,"href",qe+"/tree"),A(m,"class","mb-2"),A(h,"class","themed"),A(R,"href","https://github.com/Vilsol/timeless-jewels"),A(R,"target","_blank"),A(R,"rel","noopener"),A(w,"class","text-right text-orange-500"),A(a,"class","flex flex-col justify-between w-1/3"),A(e,"class","py-10 flex flex-row justify-center w-screen h-screen")},m(O,H){T(O,e,H),d(e,a),d(a,l),d(l,i),d(i,s),d(l,t),d(l,o),d(o,r),d(r,v),d(l,c),d(l,h),d(h,m),d(m,f),d(h,u),W(_,h,null),d(h,V),j&&j.m(h,null),d(a,b),d(a,w),d(w,R),d(R,y),M=!0},p(O,[H]){const L={};!k&&H&1&&(k=!0,L.value=O[0],X(()=>k=!1)),_.$set(L),O[0]?j?(j.p(O,H),H&1&&C(j,1)):(j=_e(O),j.c(),C(j,1),j.m(h,null)):j&&(Pe(),G(j,1,1,()=>{j=null}),Ie())},i(O){M||(C(_.$$.fragment,O),C(j),M=!0)},o(O){G(_.$$.fragment,O),G(j),M=!1},d(O){O&&p(e),Y(_),j&&j.d()}}}function Ce(n,e,a){let l,i;$e(n,Ve,b=>a(13,i=b));const s=i.url.searchParams,t=Object.keys(J.TimelessJewels).map(b=>({value:parseInt(b),label:J.TimelessJewels[b]}));let o=s.has("jewel")?t.find(b=>b.value==s.get("jewel")):void 0,r=s.has("conqueror")?{value:s.get("conqueror"),label:s.get("conqueror")}:void 0;const v=Object.values(J.PassiveSkills).map(b=>({value:b.Index,label:b.Name+" ("+b.ID+")"}));let c=s.has("passive_skill")?v.find(b=>b.value==s.get("passive_skill")):void 0,h=s.has("seed")?s.get("seed"):0,m;const f=()=>{{const b={};o&&(b.jewel=o.value),r&&(b.conqueror=r.value),c&&(b.passive_skill=c.value),h&&(b.seed=h);const w=Object.keys(b).map(R=>R+"="+encodeURIComponent(b[R])).join("&");Oe(i.url.pathname+"?"+w)}};function u(b){o=b,a(0,o)}function _(b){r=b,a(1,r)}function k(b){c=b,a(2,c)}function V(){h=Ee(this.value),a(3,h)}return n.$$.update=()=>{n.$$.dirty&1&&a(5,l=o?Object.keys(J.TimelessJewelConquerors[o.value]).map(b=>({value:b,label:b})):[]),n.$$.dirty&15&&c&&h&&o&&r&&a(4,m=Me.Calculate(c.value,typeof h=="string"?parseInt(h):h,o.value,r.value))},[o,r,c,h,m,l,t,v,f,u,_,k,V]}class Be extends Te{constructor(e){super(),je(this,e,Ce,Ne,Je,{})}}export{Be as default};
