import Vue from 'vue'
import Vuex from 'vuex'
import mixin from './mixin'

Vue.mixin(mixin)
Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    markdown: ''
  },
  getters: {
    markdown: state => {
      return state.markdown
    }
  },
  mutations: {
    markdown (state, data) {
      state.markdown = data
    }
  },
  actions: {

  }
})
