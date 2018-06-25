import "babel-polyfill";
import './styles/common.less';

import Vue from 'vue';
import VueRouter from 'vue-router';
import VueResource from 'vue-resource';
import store from './store'; 
import VueMaterial from 'vue-material';
import Toasted from 'vue-toasted';
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css'

/*Pages*/
import Beta from './pages/beta/';
import Member from './pages/member/';
import Login from './pages/member/login/';
import Create from './pages/member/create/';
import Forgot from './pages/member/forgot/';
/*End pages*/
   
import * as types from './constants/vuex-types';    
import util from './utils';  
import userApi from './api/user';

Vue.use(VueRouter); 
Vue.use(VueResource);  
Vue.use(VueMaterial);
Vue.use(Toasted)

let router = new VueRouter({
  mode: 'hash',
  routes: [
    { path: '*', name: 'Beta', component: Beta }, 
    { path: '/', name: 'Member', component: Member,
      children:[
        {path: '/', name: 'Login', component: Login},
        {path: 'login', name: 'Login', component: Login},
        {path: 'create', name: 'Create', component: Create},
        {path: 'forgot', name: 'Forgot', component: Forgot},
      ]
    }
  ]
});

new Vue({
  router,
  template: `<router-view></router-view>`,
  store
}).$mount('#app');


window.$ = $;
window.onerror = function(messageOrEvent, source, lineno, colno, error) { 
  console.error('global error handler:',messageOrEvent, source, lineno, colno, error)
}