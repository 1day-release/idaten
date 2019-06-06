<template>
  <div>
    <markdown-editor
      v-model="markdown"
      class="editor"
      @input="editorInput"
      @click.native="activePageNumber"
      @keydown.native="activePageNumber"
      ref="simplemde"
      :configs="mdeConfigs"
    />
    <div class="popups">
      <PopupText
        id="popup-logout"
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
      markdown: this.$store.getters.markdown,
      mdeConfigs: {
        toolbar: [
          'bold', 'italic', 'heading', '|', 'code', 'quote', 'unordered-list', 'ordered-list', '|', 'link', 'table'
        ]
      }
    }
  },
  methods: {
    editorInput () {
      this.$store.dispatch('updateMarkdown', this.markdown)
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
      border: none;
      background-color: transparent;
    }

    .editor-toolbar {
      flex: 0 0 auto;
      display: flex;
      padding: 2px 20px 5px 90px;

      &::before,
      &::after {
        content: none;
      }

      a {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 26px;
        height: 26px;
        font-size: 1.2rem;
        transition-duration: $duration / 2;

        &::before {
          line-height: 1.0;
        }
      }

      i {
        &.separator {
          margin: 0 10px;
          font-size: 1.0rem;
        }
      }

    }

    .editor-statusbar {
      display: none;
    }

    .CodeMirror {
      padding-right: 20px;
      padding-left: 0;
      font-size: 12px;

      &-code {
        counter-reset: line;
      }

      &-line {
        display: flex;
        align-items: flex-start;
        counter-increment: line;

        &::before {
          flex: 0 0 auto;
          content: counter(line);
          display: flex;
          justify-content: flex-end;
          align-items: center;
          width: 3.0em;
          height: 20px;
          padding-right: 0.75em;
          padding-left: 40px;
          margin-right: 0.75em;
          text-align: right;
          opacity: 0.5;
          border-top-right-radius: 10px;
          border-bottom-right-radius: 10px;
        }

        >span {
          display: block;
          height: auto;
          min-height: 20px;
        }

        .cm-header-1,
        .cm-header-2,
        .cm-header-3 {
          line-height: 1.5;
        }

        &.is-active {
          &::before {
            background-color: rgba(#fff, 0.2);
          }
        }
      }
    }
  }
</style>

<style scoped lang="scss">
  .editor {
    display: flex;
    flex-direction: column;
    overflow: hidden;
    width: 100%;
    height: 100%;
    padding: 0;
    border: none;
    margin: auto;
    line-height: 1.8;
    font-family: $ff-mono;
    color: rgba( map-get($color-brand, "text-white"), 0.9 );
    -webkit-font-smoothing: subpixel-antialiased;
    background-color: transparent;
    resize: none;
    outline: none;

    // @include scrollBarRight(7px, 7px, rgba(#fff, 0.3), rgba(#fff, 0.1));
  }

  .popups {
    position: relative;
    flex: 0 0 auto;
    margin-top: 20px;
  }

  .popup {
    position: relative;
  }

  #popup-logout {
    display: none;
    white-space: nowrap;
    margin: 0 auto;

    .l-container.is-logout & {
      top: 0;
      display: flex;
      opacity: 1;
    }
  }
</style>
