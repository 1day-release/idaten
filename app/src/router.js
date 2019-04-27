import Vue from 'vue'
import Router from 'vue-router'
import Edit from './views/Edit.vue'
import Presentation from './views/Presentation.vue'
import TestSlideComponent from './views/TestSlideComponent.vue'

Vue.use(Router)

export default new Router({
  // mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'Edit',
      component: Edit
    },
    {
      path: '/presentation',
      redirect: {
        path: '/presentation',
        params: {
          pageNumber: 1
        }
      }
    },
    {
      path: '/presentation/:pageNumber',
      name: 'Presentation',
      component: Presentation
    },
    {
      path: '/test-slide-component',
      name: 'TestSlideComponent',
      component: TestSlideComponent
    }
  ]
})
