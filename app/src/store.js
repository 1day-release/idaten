import Vue from 'vue'
import Vuex from 'vuex'
import mixin from './mixin'

import SlideModel from '@/models/slide'
import IdatenCore from 'idaten-core'

const slideModel = new SlideModel()
const core = new IdatenCore()

Vue.mixin(mixin)
Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    slides: [],
    slide: {
      markdown: ''
    },
    slideId: '',
    activePageNumber: -1,
    userSlideListState: false,
    statusMessage: '更新しました',
    statusProcessing: false
  },
  getters: {
    slides: state => {
      return state.slides
    },
    markdown: state => {
      return state.slide.markdown
    },
    slideId: state => {
      return state.slideId
    },
    activePageNumber: state => {
      return state.activePageNumber
    },
    userSlideListState: state => {
      return state.userSlideListState
    },
    statusMessage: state => {
      return state.statusMessage
    },
    statusProcessing: state => {
      return state.statusProcessing
    }
  },
  mutations: {
    slides (state, data) {
      state.slides = data
    },
    slide (state, data) {
      state.slide = data
      state.slide.markdown = data.markdown
    },
    slideId (state, data) {
      state.slideId = data
    },
    markdown (state, data) {
      state.slide.markdown = data
      state.slide.cover = core.splitPages(data)[0]
    },
    activePageNumber (state, data) {
      state.activePageNumber = data
    },
    userSlideListState (state, data) {
      state.userSlideListState = data
    },
    statusMessage (state, data) {
      state.statusMessage = data
    },
    statusProcessing (state, data) {
      state.statusProcessing = data
    }
  },
  actions: {
    async reloadList ({ commit }) {
      const slides = await slideModel.list()
      commit('slides', slides)
    },
    async setSlideId ({ commit }, slideId) {
      const slide = await slideModel.get(slideId)
      commit('slideId', slideId)
      commit('slide', slide)
      commit('markdown', slide.markdown)
    },
    async updateMarkdown ({ commit, state }, markdown) {
      commit('markdown', markdown)
      await slideModel.update(state.slideId, state.slide)
    },
    async createSlide ({ commit }) {
      const slideId = await slideModel.create()
      return slideId
    },
    async deleteSlide ({ commit }, slideId) {
      await slideModel.delete(slideId)
    },
    selectFirstSlide ({ commit, state, dispatch }) {
      const slideId = state.slides[0].id
      dispatch('setSlideId', slideId)
      return slideId
      // await slideModel.delete(slideId)
    },
    setStatus ({ commit, state, dispatch }, message) {
      if (message) {
        commit('statusProcessing', true)
        commit('statusMessage', message)
      } else {
        commit('statusProcessing', false)
        commit('statusMessage', '更新しました')
      }
    }
  }
})
