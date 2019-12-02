<template>
  <div
    class="l-container"
    :class="{'is-logout': isLogout}"
  >
    <Header />
    <Main>
      <div
        class="l-user-slide js-slidelist-container"
        :class="{'is-active': isActive}"
      >
        <UserSlideList />
      </div>
      <div class="l-editor">
        <Editor class="editor-contents" />
      </div>
      <div class="l-preview">
        <div class="preview-contents">
          <Preview :active-page-number="activePageNumber" />
        </div>
        <div class="preview-bg">
          <PreviewBg1 />
          <PreviewBg2 />
        </div>
      </div>
    </Main>
  </div>
</template>

<script>
import firebase from 'firebase/app'

import Header from '@/components/Header.vue'
import Main from '@/components/Main.vue'
import UserSlideList from '@/components/UserSlideList.vue'
import Editor from '@/components/Editor.vue'
import Preview from '@/components/Preview.vue'
import PreviewBg1 from '@/assets/preview-bg1.svg'
import PreviewBg2 from '@/assets/preview-bg2.svg'

export default {
  name: 'Edit',
  components: {
    Header,
    Main,
    UserSlideList,
    Editor,
    Preview,
    PreviewBg1,
    PreviewBg2
  },
  data () {
    return {
      isLogout: true
    }
  },
  created () {
    this.$store.dispatch('setStatus', 'ログイン中')
    firebase.auth().onAuthStateChanged(async (user) => {
      if (!user) {
        this.$router.push('/0')
      } else {
        this.isLogout = false
        if (this.$route.params.slideId === '0') {
          await this.$store.dispatch('reloadList')
          let slideId = ''
          if (this.$store.getters.slides.length === 0) {
            slideId = await this.$store.dispatch('createSlide')
          } else {
            slideId = this.$store.getters.slides[0].id
          }
          this.$store.dispatch('setSlideId', slideId)
          this.$router.push('/' + slideId)
        }
      }
      this.$store.dispatch('setStatus', '')
    })

    // this.$store.commit('slideId', this.$route.params.slideId)
    this.$store.dispatch('setSlideId', this.$route.params.slideId)
    window.pthis = this
  },
  computed: {
    isActive () {
      return this.$store.getters.userSlideListState
    },
    activePageNumber () {
      return this.$store.getters.activePageNumber
    }
  },
  watch: {
    $route (to, from) {
      this.$store.commit('slideId', this.$route.params.slideId)
    }
  },
  methods: {
  }
}
</script>

<style scoped lang="scss">
  main {
    display: flex;
    justify-content: space-between;
  }

  .l-user-slide {
    position: absolute;
    z-index: 80;
    top: 0;
    bottom: 0;
    left: -100%;
    width: 220px;
    margin: auto 0;
    transition-duration: $duration;

    &.is-active {
      left: 0;
    }
  }

  .l-editor {
    display: flex;
    justify-content: flex-end;
    overflow: hidden;
    width: calc(100% - #{$pc-min-width * 0.35 + 20px});
    height: calc(100vh - #{$pc-header-height});
    color: map-get($color-brand, "text-white");
    background-color: map-get($color-brand, "main");
  }

  .editor-contents {
    flex: 0 0 auto;
    position: relative;
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    overflow: hidden;
    width: 100%;
    height: 100%;
    padding: 20px 40px;
    padding-left: 0;
  }

  .l-preview {
    position: relative;
    overflow: hidden;
    width: $pc-min-width * 0.35 + 20px;
    height: calc(100vh - #{$pc-header-height});
    background-color: map-get($color-brand, "sub-dark");
  }

  .preview {
    &-contents {
      position: relative;
      z-index: 1;
      width: 100%;
      max-width: $pc-min-width * 0.35 + 20px;
      height: inherit;
    }

    &-bg {
      position: absolute;
      top: 0;
      right: 0;
      bottom: 0;
      left: 0;
      margin: auto;

      svg {
        position: absolute;
        width: 100%;

        &:nth-child(1) {
          top: 0;
          right: 0;
        }

        &:nth-child(2) {
          bottom: 0;
          left: 0;
        }
      }
    }
  }
</style>
