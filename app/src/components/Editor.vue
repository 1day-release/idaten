<template>
  <div>
    <markdown-editor
      v-model="markdown"
      class="editor"
      @input="editorInput"
      @click.native="activePageNumber"
      @keydown.native="activePageNumber"
      ref="simplemde"
    />
    <div class="popups">
      <PopupText
        id="popups-logout"
        class="popup"
      >
        デモモードです。ログインをしないと、「共有」「保存」などの機能を使えません。ログインしてください。
      </PopupText>
    </div>
  </div>
</template>

<script>
import MarkdownEditor from 'vue-simplemde/src/markdown-editor'
import PopupText from '@/components/PopupText.vue'
import IdatenCore from 'idaten-core'
const core = new IdatenCore()

export default {
  name: 'Editor',
  components: {
    MarkdownEditor,
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
    },
    activePageNumber () {
      const codemirror = this.$refs.simplemde.simplemde.codemirror
      const activeLineNumber = codemirror.getCursor('start').line + 1
      const lineNumbers = core.getPageLineNumbers(this.markdown)
      let activePageNumber = 1
      lineNumbers.forEach((number, index) => {
        if (number < activeLineNumber) {
          activePageNumber = index + 2
        }
      })
      this.$store.commit('activePageNumber', activePageNumber)
    }
  },
  watch: {
    '$store.getters.markdown' () {
      this.markdown = this.$store.getters.markdown
    }
  },
  mounted () {
    this.$store.commit('activePageNumber', 1)
  }
}
</script>

<style lang="scss">
  @import '../simplemde-theme-dark.min.css';
  .editor {
    .editor-toolbar, .CodeMirror {
      background-color: transparent;
    }
    .CodeMirror {
      height: 75vh; // 暫定
    }
  }
</style>

<style scoped lang="scss">
  $width: 100%;
  $max-width: $pc-min-width * 0.65;

  .editor {
    position: absolute;
    top: 40px;
    right: 40px;
    bottom: 70px;
    width: $width;
    max-width: $max-width;
    padding: 0;
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
