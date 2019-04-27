<template>
  <div>
    <textarea class="editor" v-model="markdown" @input="editorInput" />
    <div class="popups">
      <PopupText id="popups-logout" class="popup">デモモードです。ログインをしないと、「共有」「保存」などの機能を使えません。<a href="#">ログイン</a>してください。</PopupText>
    </div>
  </div>
</template>

<script>
import PopupText from '@/components/PopupText.vue'
import IdatenCore from 'idaten-core'
const core = new IdatenCore()

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
    },
    setCursorPositionEvent () {
      const editor = this.$el.querySelector('.editor')
      const setPageNumber = () => {
        setTimeout(() => {
          const cursorStart = editor.selectionStart
          const corsorPosition = this.markdown.slice(0, cursorStart).split('\n').length || -1
          const lineNumbers = core.getPageLineNumbers(this.markdown)
          let activePageNumber = 1
          lineNumbers.forEach((number, index) => {
            if (number < corsorPosition) {
              activePageNumber = index + 2
            }
          })
          this.$store.commit('activePageNumber', activePageNumber)
        }, 50)
      }
      editor.addEventListener('keydown', setPageNumber, false)
      editor.addEventListener('click', setPageNumber, false)
      this.$store.commit('activePageNumber', 1)
    }
  },
  mounted () {
    this.setCursorPositionEvent()
  }
}
</script>

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
