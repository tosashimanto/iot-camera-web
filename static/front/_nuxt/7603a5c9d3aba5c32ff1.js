(window.webpackJsonp=window.webpackJsonp||[]).push([[10],{319:function(e,t,r){var a=r(337);"string"==typeof a&&(a=[[e.i,a,""]]),a.locals&&(e.exports=a.locals);(0,r(40).default)("268d3a3a",a,!0,{})},320:function(e,t,r){var a=r(339);"string"==typeof a&&(a=[[e.i,a,""]]),a.locals&&(e.exports=a.locals);(0,r(40).default)("35155038",a,!0,{})},336:function(e,t,r){"use strict";var a=r(319);r.n(a).a},337:function(e,t,r){(e.exports=r(39)(!1)).push([e.i,".VueToNuxtLogo{display:inline-block;-webkit-animation:turn 2s linear 1s forwards;animation:turn 2s linear 1s forwards;-webkit-transform:rotateX(180deg);transform:rotateX(180deg);position:relative;overflow:hidden;height:180px;width:245px}.Triangle{position:absolute;top:0;left:0;width:0;height:0}.Triangle--one{border-left:105px solid transparent;border-right:105px solid transparent;border-bottom:180px solid #41b883}.Triangle--two{top:30px;border-left:87.5px solid transparent;border-right:87.5px solid transparent;border-bottom:150px solid #3b8070}.Triangle--three,.Triangle--two{left:35px;-webkit-animation:goright .5s linear 3.5s forwards;animation:goright .5s linear 3.5s forwards}.Triangle--three{top:60px;border-left:70px solid transparent;border-right:70px solid transparent;border-bottom:120px solid #35495e}.Triangle--four{top:120px;left:70px;-webkit-animation:godown .5s linear 3s forwards;animation:godown .5s linear 3s forwards;border-left:35px solid transparent;border-right:35px solid transparent;border-bottom:60px solid #fff}@-webkit-keyframes turn{to{-webkit-transform:rotateX(0deg);transform:rotateX(0deg)}}@keyframes turn{to{-webkit-transform:rotateX(0deg);transform:rotateX(0deg)}}@-webkit-keyframes godown{to{top:180px}}@keyframes godown{to{top:180px}}@-webkit-keyframes goright{to{left:70px}}@keyframes goright{to{left:70px}}",""])},338:function(e,t,r){"use strict";var a=r(320);r.n(a).a},339:function(e,t,r){(e.exports=r(39)(!1)).push([e.i,".el-footer,.el-header{background-color:#b3c0d1;color:#333;text-align:center;line-height:60px}.el-aside{background-color:#d3dce6;line-height:90vh}.el-aside,.el-main{color:#333;text-align:center}.el-main{background-color:#e9eef3;line-height:160px}body>.el-container{margin-bottom:40px}",""])},341:function(e,t,r){"use strict";r.r(t);r(20);var a=r(1),n=r.n(a),o=(r(336),r(5)),i=Object(o.a)({},function(){this.$createElement;this._self._c;return this._m(0)},[function(){var e=this.$createElement,t=this._self._c||e;return t("div",{staticClass:"VueToNuxtLogo"},[t("div",{staticClass:"Triangle Triangle--two"}),this._v(" "),t("div",{staticClass:"Triangle Triangle--one"}),this._v(" "),t("div",{staticClass:"Triangle Triangle--three"}),this._v(" "),t("div",{staticClass:"Triangle Triangle--four"})])}],!1,null,null,null).exports,l=(r(52),{data:function(){return{tree_data:[{parent_id:"",label:"Level one 1",value:"",children:[{parent_id:"",label:"Level two 1-1",value:""},{parent_id:"",label:"Level two 1-2",value:""}]},{parent_id:"",label:"Level one 2",value:"",children:[{parent_id:"",label:"Level two 2-1",value:""},{parent_id:"",label:"Level two 2-2",value:""}]}],defaultProps:{children:"children",label:"label"}}},asyncData:function(){var e=n()(regeneratorRuntime.mark(function e(t){var r;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return r=t.store,t.route,e.next=3,r.dispatch("images/getImageList");case 3:return e.next=5,r.dispatch("d3_map/getMapJSON");case 5:case"end":return e.stop()}},e,this)}));return function(t){return e.apply(this,arguments)}}(),components:{Logo:i},methods:{getTabeldata:function(){return[{date:"2016-05-03",name:"Tom",address:"No. 189, Grove St, Los Angeles"},{date:"2016-05-02",name:"Tom",address:"No. 189, Grove St, Los Angeles"},{date:"2016-05-04",name:"Tom",address:"No. 189, Grove St, Los Angeles"},{date:"2016-05-01",name:"Tom",address:"No. 189, Grove St, Los Angeles"}].map(function(e){return e})},handleClickImages:function(){var e=n()(regeneratorRuntime.mark(function e(){return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:this.$router.push("/images/");case 1:case"end":return e.stop()}},e,this)}));return function(){return e.apply(this,arguments)}}(),handleClickD3map:function(){var e=n()(regeneratorRuntime.mark(function e(){return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:this.$router.push("/d3_map/");case 1:case"end":return e.stop()}},e,this)}));return function(){return e.apply(this,arguments)}}(),handleNodeClick:function(e){console.log(e)},handleClickD3graph:function(){console.log("handleClickD3graph"),this.$router.push("/d3_graph/")},handleClickwebSocket:function(){console.log("handleClickwebSocket"),this.$router.push("/websocket/")},handleClickUpload:function(){console.log("handleClickUpload"),this.$router.push("/file-upload/")}}}),s=(r(338),Object(o.a)(l,function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("section",{staticClass:"container posts-page"},[r("el-card",{staticStyle:{flex:"1"}},[r("el-container",[r("el-header",[e._v("\n        Header\n      ")]),e._v(" "),r("el-container",[r("el-aside",{attrs:{width:"200px"}},[r("el-tree",{attrs:{data:e.tree_data,props:e.defaultProps},on:{"node-click":e.handleNodeClick}})],1),e._v(" "),r("el-container",[r("el-main",[r("el-table",{staticStyle:{width:"100%"},attrs:{data:e.getTabeldata(),stripe:""}},[r("el-table-column",{attrs:{prop:"date",label:"Date",width:"180"}}),e._v(" "),r("el-table-column",{attrs:{prop:"name",label:"Name",width:"180"}}),e._v(" "),r("el-table-column",{attrs:{prop:"address",label:"Address"}})],1),e._v(" "),r("div",{staticClass:"text-right"},[r("el-button",{attrs:{type:"primary"},on:{click:e.handleClickImages}},[e._v("Test")]),e._v(" "),r("el-button",{attrs:{type:"primary"},on:{click:e.handleClickD3map}},[e._v("D3 map")]),e._v(" "),r("el-button",{attrs:{type:"primary"},on:{click:e.handleClickD3graph}},[e._v("D3 graph")]),e._v(" "),r("el-button",{attrs:{type:"primary"},on:{click:e.handleClickwebSocket}},[e._v("WebSocket")]),e._v(" "),r("el-button",{attrs:{type:"primary"},on:{click:e.handleClickUpload}},[e._v("upload")])],1)],1)],1)],1),e._v(" "),r("el-footer",[e._v("\n        Footer\n      ")])],1)],1)],1)},[],!1,null,null,null));t.default=s.exports}}]);