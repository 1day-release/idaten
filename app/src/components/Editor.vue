<template>
  <div>
    <textarea class="editor" v-model="markdown" @input="editorInput">
    </textarea>
    <div class="popups">
      <PopupText id="popups-logout" class="popup">デモモードです。ログインをしないと、「共有」「保存」などの機能を使えません。<a href="#">ログイン</a>してください。</PopupText>
    </div>
  </div>
</template>

<script>
import PopupText from '@/components/PopupText.vue'

export default {
  name: 'Editor',
  components: {
    PopupText
  },
  props: {
  },
  data () {
    return {
      markdown: this.$store.getters.markdown
    }
  },
  methods: {
    editorInput () {
      this.$store.commit('markdown', this.markdown)
    }
  },
  created () {
  }
}
</script>

<style scoped lang="scss">
  $width: 100%;
  $max-width: $pc-min-width * 0.65 + 20px;

  .editor {
    box-sizing: border-box;
    position: absolute;
    top: 40px;
    right: 40px;
    bottom: 70px;
    width: $width;
    max-width: $max-width;
    padding-right: 20px;
    border: none;
    margin: auto;
    line-height: 1.8;
    font-family: $ff-mono;
    color: rgba( map-get($color-brand, "text-white"), 0.9 );
    -webkit-font-smoothing: subpixel-antialiased;
    background-color: transparent;
    resize: none;
    outline: none;

    @include scrollBarRight(7px, 7px, rgba(#fff, 0.3), rgba(#fff, 0.1));
  }

  .popups {
    position: absolute;
    right: 40px;
    bottom: 20px;
    width: $width;
    max-width: $max-width;
    height: 30px;
  }

  .popup {
    position: absolute;
    right: 0;
  }
</style>
