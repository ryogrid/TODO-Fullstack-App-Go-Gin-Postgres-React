(window.webpackJsonp=window.webpackJsonp||[]).push([[0],[,,,,,,,,,function(e,t,n){e.exports=n(18)},,,,,function(e,t,n){},function(e,t,n){},function(e,t,n){},function(e,t,n){},function(e,t,n){"use strict";n.r(t);var a=n(0),o=n.n(a),i=n(7),c=n.n(i),r=n(1),s=n(2),l=n(4),u=n(3),m=n(5),d=(n(14),function(e){function t(){return Object(r.a)(this,t),Object(l.a)(this,Object(u.a)(t).apply(this,arguments))}return Object(m.a)(t,e),Object(s.a)(t,[{key:"render",value:function(){return o.a.createElement("div",{className:"Header"}," TODO List ")}}]),t}(o.a.Component)),h=(n(15),function(e){function t(e){var n;return Object(r.a)(this,t),(n=Object(l.a)(this,Object(u.a)(t).call(this,e))).bkend_url_base=null,n.addItem=function(e){"Enter"===e.key&&fetch(n.bkend_url_base+"/item/create/".concat(e.target.value)).then(window.location.reload())},"localhost"==window.location.hostname?n.bkend_url_base="http://localhost:8088":n.bkend_url_base="http://"+window.location.hostname+":8088",n}return Object(m.a)(t,e),Object(s.a)(t,[{key:"render",value:function(){return o.a.createElement("div",{className:"AddBar"},o.a.createElement("input",{className:"AddBar-Text",type:"text",placeholder:"Enter TODO Item",onKeyDown:this.addItem}))}}]),t}(o.a.Component)),b=n(8),f=(n(16),function(e){function t(e){var n;return Object(r.a)(this,t),(n=Object(l.a)(this,Object(u.a)(t).call(this,e))).bkend_url_base=null,n.state={items:[]},"localhost"==window.location.hostname?n.bkend_url_base="http://localhost:8088":n.bkend_url_base="http://"+window.location.hostname+":8088",n}return Object(m.a)(t,e),Object(s.a)(t,[{key:"removeItem",value:function(e){fetch(this.bkend_url_base+"/item/delete/".concat(e)).then(this.setState({items:this.state.items.filter(function(t){return t.id!==e})}))}},{key:"toggleDone",value:function(e){var t=Object(b.a)(this.state.items),n=t.find(function(t){return t.id===e});n.done=!n.done,fetch(this.bkend_url_base+"/item/update/".concat(e,"/").concat(n.done)).then(this.setState({items:t}))}},{key:"isDone",value:function(e){return e?"Done":"Not Done"}},{key:"createItem",value:function(e){var t=this;return o.a.createElement("div",{className:"ListItem",key:e.id,id:e.id},o.a.createElement("div",{className:"Title"},o.a.createElement("div",{className:"RemoveItem",onClick:function(){return t.removeItem(e.id)}},"X"),e.item),o.a.createElement("div",{className:"Status",onClick:function(){return t.toggleDone(e.id)}},this.isDone(e.done)))}},{key:"componentDidMount",value:function(){var e=this;fetch(this.bkend_url_base+"/items").then(function(e){return e.json()}).then(function(t){return e.setState({items:t.items})})}},{key:"render",value:function(){var e=this,t=this.state.items;return o.a.createElement("div",{className:"TodoList"},o.a.createElement("div",{className:"List"},t.map(function(t){return e.createItem(t)})))}}]),t}(o.a.Component)),v=(n(17),function(e){function t(){return Object(r.a)(this,t),Object(l.a)(this,Object(u.a)(t).apply(this,arguments))}return Object(m.a)(t,e),Object(s.a)(t,[{key:"render",value:function(){return o.a.createElement("div",{className:"App"},o.a.createElement(d,null),o.a.createElement(h,null),o.a.createElement(f,null))}}]),t}(o.a.Component));Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));c.a.render(o.a.createElement(v,null),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then(function(e){e.unregister()})}],[[9,1,2]]]);
//# sourceMappingURL=main.c1c1a39b.chunk.js.map