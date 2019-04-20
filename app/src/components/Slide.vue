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
  $ratio: 0.618;

  $width: 1112px;
  $height: $width * $ratio;
  $padding: 30px;

  $innerWidth: $width - $padding * 2;
  $innerHeight: $height - $padding * 2;

  $base-fz: 18px;

  $fz: (
           "pager": 12px,

           "title": 40px,
       "sub-title": 20px,
              "to": 16px,
            "date": 16px,
            "from": 16px,

     "cover-title": 12px,
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
    position: relative;
    display: flex;
    flex-direction: column;
    width: 100%;
    height: auto;
    // min-height: 680px / $width * 100%;
    color: #0f002b;
    font-size: map-get($fz, "p");
    letter-spacing: 0.05em;
    background: #fff;

    // Layout
    // ==============================
    .page {
      &-header,
      &-body,
      &-footer {
        box-sizing: border-box;
        width: 100%;
      }

      &-header {
      }

      &-body {
      }

      &-footer {
      }
    }

    // Parts
    // ==============================
    * {
      user-select: none;
    }

    img {
      max-width: 100%;
    }

    .page {
      &-number {
        display: flex;
        justify-content: center;
        align-items: center;
        @include fontSizeRatio( map-get($fz, "pager") );
        color: rgba( map-get($color-brand, "main"), 0.4 );

        &::before,
        &::after {
          content: "";
          width: 0.25em;
          height: 1px;
          margin: 0 0.5em;
          background-color: rgba( map-get($color-brand, "main"), 0.4 );
        }
      }

      &-total {
        &::before {
          content: "/";
          margin: 0 0.5em;
        }
      }
    }

    &.is-type1 {
      justify-content: space-between;
      padding: 80px / $width * 100%;

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
            margin: 15px / $height * 100% 0;
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
            margin-top: 15px / $height * 100%;
            @include fontSizeRatio( map-get($fz, "from") );
          }
        }
      }
    }

    &.is-type2,
    &.is-type3 {
      justify-content: center;
      align-items: center;
      padding: (80px+20px+$padding)/$width*100% $padding/$width*100%;

      /deep/ {
        p {
          text-align: center;
          line-height: 1.5;
        }

        .slide {
          &-header {

            .cover {
              &-title {
                @include fontSizeRatio( map-get($fz, "cover-title") );
                color: rgba( map-get($color-brand, "main"), 0.4 );
              }
            }
          }
        }

        .page {
          &-header {
          }

          &-body {
            padding: 0 35px / $innerWidth * 100%;
          }

          &-footer {
            position: absolute;
            bottom: 0;
            display: flex;
            justify-content: center;
            margin: 0 auto;
            margin-bottom: 30px / $height * 100%;
          }
        }
      }
    }

    &.is-type2 {

      /deep/ {
        .slide {
          &-header {
            position: absolute;
            top: $padding / $height * 100%;
            left: $padding / $width * 100%;
          }
        }

        .page {
          &-header {
          }

          &-body {
            margin-top: 20px / $innerHeight * 100%;
          }

          &-footer {
          }
        }

        .childcover {
          &-title {
            display: flex;
            flex-direction: column;
            align-items: center;
            @include fontSizeRatio( map-get($fz, "h2") );
            font-weight: bold;
          }

          &-number {
            margin-bottom: 20px / $innerWidth * 100%;
            font-size: 50%;
            color: map-get($color-brand, "text-sub");
          }
        }

        p {
          @include fontSizeRatio( 20px );
        }
      }
    }

    &.is-type3 {
      justify-content: flex-start;
      padding-top: $padding / $width * 100%;

      p {
        @include fontSizeRatio( $base-fz );
      }

      /deep/ {
        .slide {
          &-header {
            margin-bottom: 12px / $innerWidth * 100%;
            align-self: self-start;
          }
        }

        .page {
          &-header {
            margin-bottom: 40px / $width * 100%;
            text-align: center;
          }

          &-body {
            overflow-y: auto;
            height: 100%;
          }

          &-footer {
          }
        }

        .childcover {
          &-title {
            display: flex;
            justify-content: center;
            align-items: flex-end;
            margin-bottom: 10px / $width * 100%;
            @include fontSizeRatio( map-get($fz, "section-title") );
          }

          &-number {
            margin-right: 0.3em;
            @include fontSizeRatio( 12px );
          }
        }

        .contents {
          &-title {
            line-height: 1.5;
            @include fontSizeRatio( map-get($fz, "h3") );
            font-weight: bold;
          }
        }

        ol {
          list-style-type: decimal;
        }

        ul {
          list-style-type: disc;
        }
      }
    }
  }
</style>
