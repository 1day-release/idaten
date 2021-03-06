import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import mixin from './mixin'

import './firebaseApp'

import VueAnalytics from 'vue-analytics'

import 'reset-css'
import 'source-code-pro/source-code-pro.css'

Vue.mixin(mixin)
Vue.use(VueAnalytics, {
  id: 'UA-135478557-4',
  router
})
Vue.config.productionTip = false

require('@/styles/styles.scss')

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
