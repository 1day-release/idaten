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
    slides: slideModel.list(),
    slide: {},
    slideId: '',
    activePageNumber: -1,
    userSlideListState: false
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
    }
  },
  mutations: {
    slides (state, data) {
      state.slides = slideModel.list()
    },
    markdown (state, data) {
      state.slide.markdown = data
      state.slide.cover = core.splitPages(data)[0]
      slideModel.update(state.slideId, state.slide)
    },
    slideId (state, data) {
      state.slideId = data
      state.slide = slideModel.get(data)
      state.markdown = state.slide.markdown
    },
    activePageNumber (state, data) {
      state.activePageNumber = data
    },
    userSlideListState (state, data) {
      state.userSlideListState = data
    }
  },
  actions: {
    changeStateUserSlideList ({ commit }, state) {
      commit('userSlideListState', state)
    },
    createSlide ({ commit }, state) {
      const slide = {
        cover: '# サンプルについて',
        markdown: `# サンプルについて
サブタイトル : サブタイトル
日付 : 2019-01-01
宛名 : 〇〇様
製作者 : 〇〇

## Slide1

### Slide2
1. Ordered List1
2. Ordered List2
`
      }
      slideModel.create(slide)
      commit('slides')
    }
  }
})
