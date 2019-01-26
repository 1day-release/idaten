import Vue from 'vue'
import Router from 'vue-router'
import Edit from './views/Edit.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'edit',
      component: Edit
    }
  ]
})
