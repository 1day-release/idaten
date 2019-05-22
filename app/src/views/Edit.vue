<template>
  <div class="l-container is-logout">
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
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
import Header from '@/components/Header.vue'
import Main from '@/components/Main.vue'
import UserSlideList from '@/components/UserSlideList.vue'
import Editor from '@/components/Editor.vue'
import Preview from '@/components/Preview.vue'
import PreviewBg1 from '@/assets/preview-bg1.svg'
import PreviewBg2 from '@/assets/preview-bg2.svg'
import queryString from 'query-string'

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
  created () {
    const parsed = queryString.parse(location.search)
    if (parsed.state === 'idaten' && parsed.code) {
      fetch('/api/auth/login', {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          code: parsed.code
        })
      }).then(response => response.json()).then(result => {
        localStorage.setItem('idaten.access-token', result.access_token)
        location.reload()
      })
    }
    this.$store.commit('slideId', this.$route.params.slideId)
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
