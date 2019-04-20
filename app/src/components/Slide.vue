<template>
  <div>
    <div v-if="html" v-html="html"></div>
    <div v-else v-html="markedHtml"></div>
  </div>
</template>

<script>
import IdatenCore from 'idaten-core'
const core = new IdatenCore()

export default {
  name: 'Slide',
  data () {
    return {
    }
  },
  props: {
    markdown: {
      type: String,
      default: ''
    },
    pageNumber: {
      type: Number,
      default: 0
    },
    html: {
      type: String,
      default: ''
    },
    width: {
      type: Number,
      default: null
    },
    maxWidth: {
      type: Number,
      default: null
    },
    maxHeight: {
      type: Number,
      default: null
    }
  },
  mounted () {
    this.computePageStyles()
  },
  updated () {
    this.computePageStyles()
  },
  methods: {
    computePageStyles () {
      let styles = {}

      if (this.width) {
        styles['font-size'] = Math.round(this.width / 61.8)
        styles['min-height'] = Math.round(this.width * 0.618)
      } else if (this.maxWidth && this.maxHeight) {
        const tempMaxWidth = Math.round(this.maxHeight / 0.618)
        const tempMaxHeight = Math.round(this.maxWidth * 0.618)
        if (tempMaxWidth > this.maxWidth) {
          styles['width'] = this.maxWidth
          styles['height'] = tempMaxHeight
        } else {
          styles['width'] = tempMaxWidth
          styles['height'] = this.maxHeight
        }
        styles['font-size'] = Math.round(styles['width'] / 61.8)
      }

      this.$nextTick(() => {
        const pageElement = this.$el.querySelector('.page')
        if (!pageElement) return
        const style = Object.entries(styles).map(style => {
          return `${style[0]}:${style[1]}px`
        }).join(';') + ';'
        pageElement.setAttribute('style', style)
      })
      return styles
    }
  },
  computed: {
    markedHtml () {
      return core.get(this.markdown, this.pageNumber)
    }
  }
}
</script>

<style scoped lang="scss">
  $fz: (
               "p": 18px,
           "pager": 12px,

           "title": 40px,
       "sub-title": 20px,
              "to": 16px,
            "date": 16px,
            "name": 16px,

     "slide-title": 12px,
              "h2": 40px,

   "section-title": 14px,
              "h3": 30px,
              "h4": 22px
  );

  /deep/ .page {
    box-sizing: border-box;
    width: 100%;
    height: auto;
    min-height: 680px / 1110px * 100%;
    color: #0f002b;
    font-size: 18px;
    background: #fff;

    * {
      user-select: none;
    }

    &.is-type1  /deep/ {
      .to {
      }

      h1 {
      }

      .subtitle {
      }

      .date {
      }

      .from {
      }
      // padding: 5% 5%;
      // box-sizing: border-box;

      // p:nth-of-type(1) {
      //   font-size: 16px;
      // }

      // h1 {
      //   font-size: 30px;
      //   margin-top: 25%;
      //   margin-bottom: 2%;
      // }

      // p:nth-of-type(2) {
      //   margin-bottom: 25%;
      // }

      // p:nth-of-type(3) {
      //   color: #cbcdd6;
      //   margin-bottom: 1%;
      // }

      // p:nth-of-type(4) {
      //   color: #cbcdd6;
      // }
    }

    &.is-type2  /deep/ {
      h2 {
        margin-top: 30%;
        font-size: 30px;
        color: #000c36;
        text-align: center;
      }
    }

    &.is-type3  /deep/ {
      padding: 10% 10%;
      box-sizing: border-box;

      h2 {
        font-size: 30px;
        color: #000c36;
        text-align: center;
      }

      h3 {
        margin-bottom: 5%;
        font-size: 30px;
        color: #000c36;
        text-align: center;
      }

      ol {
        list-style-type: decimal;
      }

      ul {
        list-style-type: disc;
      }
    }
  }
</style>
