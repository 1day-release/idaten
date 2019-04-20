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
    this.$watch('pageStyles', () => {}) // For detect props changing
    this.$watch('markdown', () => { this.computePageStyles() })
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
    pageStyles () {
      return this.computePageStyles()
    },
    markedHtml () {
      return core.get(this.markdown, this.pageNumber)
    }
  }
}
</script>

<style scoped lang="scss">
  $size: 1112px;

  $base-fz: 18px;

  $fz: (
           "pager": 12px,

           "title": 40px,
       "sub-title": 20px,
              "to": 16px,
            "date": 16px,
            "from": 16px,

     "slide-title": 12px,
              "h2": 40px,

   "section-title": 14px,
              "h3": 30px,
              "h4": 22px
  );

  @mixin fontSizeRatio($font: 18px) {
    font-size: $font / $base-fz * 100%;
  }

  /deep/ .page {
    box-sizing: border-box;
    width: 100%;
    height: auto;
    min-height: 680px / $size * 100%;
    color: #0f002b;
    font-size: map-get($fz, "p");
    letter-spacing: 0.05em;
    background: #fff;

    * {
      user-select: none;
    }

    &.is-type1 {
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      padding: 80px / $size * 100%;

      /deep/ {

        .page {
          &-header {
          }

          &-body {
            font-weight: bold;
          }

          &-footer {
            color: rgba( map-get($color-brand, "main"), 0.4);
          }
        }

        .cover {
          &-to {
            @include fontSizeRatio( map-get($fz, "to") );
          }

          &-title {
            margin: 15px 0;
            @include fontSizeRatio( map-get($fz, "title") );
          }

          &-subtitle {
            @include fontSizeRatio( map-get($fz, "sub-title") );
            color: map-get($color-brand, "text-sub");
          }

          &-date {
            @include fontSizeRatio( map-get($fz, "date") );
          }

          &-from {
            margin-top: 15px;
            @include fontSizeRatio( map-get($fz, "from") );
          }
        }
      }
    }

    &.is-type2  /deep/ {
    }

    &.is-type3  /deep/ {

      ol {
        list-style-type: decimal;
      }

      ul {
        list-style-type: disc;
      }
    }
  }
</style>