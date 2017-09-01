import Vue from 'vue'
import Vuetify from 'vuetify'

import App from '@/App.vue'
import router from './router'

import '../node_modules/vuetify/dist/vuetify.min.css'

Vue.use(Vuetify)

Vue.config.productionTip = false

/* eslint-disable no-new */
let v = new Vue({
  el: '#app',
  router,
  render: h => h(App)
})
