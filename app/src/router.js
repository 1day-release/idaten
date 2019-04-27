import Vue from 'vue'
import Router from 'vue-router'
import Edit from './views/Edit.vue'
import Presentation from './views/Presentation.vue'
import TestSlideComponent from './views/TestSlideComponent.vue'

import SlideModel from '@/models/slide'
const slideModel = new SlideModel()

Vue.use(Router)

export default new Router({
  // mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      redirect: {
        path: '/' + slideModel.list()[0].slide_id
      }
    },
    {
      path: '/:slideId',
      name: 'Edit',
      component: Edit
    },
    {
      path: '/presentation',
      redirect: {
        path: '/presentation/' + slideModel.list()[0].slide_id + '/1'
      }
    },
    {
      path: '/presentation/:slideId/:pageNumber',
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
