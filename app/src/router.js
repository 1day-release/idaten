import Vue from 'vue'
import Router from 'vue-router'
import Edit from './views/Edit.vue'
import Presentation from './views/Presentation.vue'
import Embed from './views/Embed.vue'
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
      path: '/:slideId',
      name: 'Edit',
      component: Edit
    },
    /*
    {
      path: '/presentation',
      redirect: {
        path: '/presentation/' + slideModel.list()[0].slide_id + '/1'
      }
    },
    */
    {
      path: '/presentation/:slideId/:pageNumber',
      name: 'Presentation',
      component: Presentation
    },
    /*
    {
      path: '/embed',
      redirect: {
        path: '/embed/' + slideModel.list()[0].slide_id + '/1'
      }
    },
    */
    {
      path: '/embed/:slideId/:pageNumber',
      name: 'Embed',
      component: Embed
    },
    {
      path: '/test-slide-component',
      name: 'TestSlideComponent',
      component: TestSlideComponent
    }
  ]
})
