<template>
  <div>
    <Header />
    <Main>
      <div class="l-user-slide js-slidelist-container"  v-bind:class="{ 'is-active': isActive}" >
        <UserSlideList />
      </div>
      <Editor class="l-editor" />
      <div class="l-preview">
        <div class="preview-contents">
          <Preview />
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
  computed: {
    isActive () {
      return this.$store.getters.userSlideListState
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
    position: relative;
    width: 65%;
    height: calc(100vh - #{$pc-header-height});
    color: map-get($color-brand, "text-white");
    background-color: map-get($color-brand, "main");
  }

  .l-preview {
    position: relative;
    overflow: hidden;
    width: 35%;
    height: calc(100vh - #{$pc-header-height});
    background-color: map-get($color-brand, "sub-dark");
  }

  .preview {
    &-contents {
      position: relative;
      z-index: 1;
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
