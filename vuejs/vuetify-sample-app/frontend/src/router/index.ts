import Vue from 'vue'
import Router, { RouteConfig } from 'vue-router'

import Hello from '@/components/Hello.vue'
import Home from '@/pages/Home.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    } as RouteConfig,
    {
      path: '/hello',
      name: 'Hello',
      component: Hello
    } as RouteConfig
  ]
})
