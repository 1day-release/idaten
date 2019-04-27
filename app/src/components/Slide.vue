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
      default: 1
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
  },
  updated () {
    this.computePageStyles()
  }
}
</script>

<style scoped lang="scss">
  $ratio: 0.618;

  $width: 1100px;
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

    // Modules
    // ==============================
    .content {
      &s {
      }

      box-sizing: border-box;
      margin: 40px / $innerWidth * 100% 0;
      text-align: center;

      &:first-child {
        margin-top: 0;
      }

      &:last-child {
        margin-bottom: 0;
      }

      &.is-type1 {
      }

      &.is-type2 {
      }
    }

    .item {
      &s {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        padding: 0 35px / $innerWidth * 100%;

        &[class*=is-column] .item p {
          text-align: left;
        }

        &.is-column2 {
          // slide 1100
          // items  940+40 980
          // item   450+40 490
          // padding 20

          padding: 0 (50px-20px) / $innerWidth * 100%;
          // margin: 0 -20px / $innerWidth * 100%;

          .item {
            width: calc(100% / 2);
            padding: 0 20px / 980px * 100%;

            h4 {
              margin-bottom: 20px / 450px * 100%;
            }

            p {
            }
          }
        }

        &.is-column3 {
          // slide 1100
          // items  940+44 984
          // item   284+44 328
          // padding 22

          padding: 0 (50px-22px) / $innerWidth * 100%;

          .item {
            width: calc(100% / 3);
            padding: 0 22px / 988px * 100%;

            h4 {
              margin-bottom: 20px / 284px * 100%;
            }

            p {
            }
          }
        }

        &.is-column4 {
          // slide 1100
          // items  940+40 980
          // item   225+20
          // padding 10

          // padding: 0 10px / $innerWidth * 100%;
          padding: 0;

          .item {
            width: calc(100% / 4);
            padding: 0 10px / 988px * 100%;

            h4 {
              margin-bottom: 20px / 225px * 100%;
            }

            p {
            }
          }
        }
      }

      box-sizing: border-box;
      flex: 0 0 auto;
      width: 100%;
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

    // Parts
    // ==============================
    * {
      user-select: none;
    }

    h4 {
      margin-bottom: 20px / $innerWidth * 100%;
      text-align: center;
      @include fontSizeRatio( map-get($fz, "h4") );
      font-weight: bold;
      color: map-get($color-brand, "text-sub");
    }

    ol,
    ul {
      text-align: left;
      line-height: 1.5;

      >li {
        position: relative;
        padding-left: 2em;
        margin-top: 15px / $innerWidth * 100%;

        &::before {
          content: "";
          position: absolute;
          left: 0;
        }
      }
    }

    ol {
      counter-reset: ol;

      >li {
        counter-increment: ol;

        &::before {
          width: 1.5em;
          text-align: right;
          content: counter(ol) ".";
        }

        &:first-child {
          margin-top: 0;
        }
      }
    }

    ul {

      >li {

        &::before {
          $size: 4px;

          content: "";
          top: 0.5em;
          display: block;
          width: $size;
          height: $size;
          margin-left: 6px;
          border-radius: 50%;
          background-color: map-get($color-brand, "sub");
        }
      }
    }

    img {
      display: block;
      max-width: 100%;
      margin: 0 auto;
    }

    em {
      font-weight: bold;
      font-style: oblique;
    }

    strong {
      font-weight: bold;
      color: map-get($color-brand, "text-sub");
    }

    a {
      color: #0089ff;

      &:hover {
        text-decoration: none;
      }
    }

    // Statement
    // ==============================
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
            top: 0;
            left: 0;
            margin: $padding / $width * 100%;
            margin-bottom: 0;
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

            p {
              @include fontSizeRatio( $base-fz );
            }
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
      }
    }
  }
</style>
