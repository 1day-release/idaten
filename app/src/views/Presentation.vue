<template>
  <div class="presentation" @click="nextSlide" @contextmenu.prevent="prevSlide">
    <div class="presentation-header">
      <p class="presentation-close">
        <a href="#">
          <span class="presentation-close-button">Esc</span>で編集ページに戻る
        </a>
      </p>
    </div>
    <div class="presentation-body">
      <Slide :markdown="slideMarkdown[nowSlideNum]" :max-width="slideMaxWidth" :max-height="slideMaxHeight"  />
    </div>
    <div class="presentation-footer">
      <ul class="presentation-pager">
        <li><a href="#1">1</a></li>
        <li><a class="is-now" href="#2">2</a></li>
        <li><a href="#3">3</a></li>
        <li><a href="#4">4</a></li>
        <li><a href="#5">5</a></li>
      </ul>
      <p class="presentation-logo">
        <a href="/">
          <BrandLogo class="is-gray" />
        </a>
      </p>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
import Slide from '@/components/Slide.vue'
import BrandLogo from '@/components/BrandLogo.vue'

export default {
  name: 'Presentation',
  components: {
    Slide,
    BrandLogo
  },
  data () {
    return {
      nowSlideNum: 0,
      markdown: '',
      slideMaxWidth: 0,
      slideMaxHeight: 0
    }
  },
  created () {
    if (this.$route.query.mdUrl) {
      console.log('md', this.$route.query.mdUrl)
      fetch(this.$route.query.mdUrl).then(request => request.text()).then((markdown) => {
        this.markdown = markdown
      })
    }
    window.addEventListener('keydown', (event) => {
      if (event.keyCode === 39) {
        this.nextSlide()
      } else if (event.keyCode === 37) {
        this.prevSlide()
      }
    })

    window.addEventListener('resize', this.calculateSlideMaxSize)
    this.calculateSlideMaxSize()
  },
  methods: {
    nextSlide () {
      if (this.nowSlideNum < this.slideMarkdown.length - 1) this.nowSlideNum++
    },
    prevSlide () {
      if (this.nowSlideNum > 0) this.nowSlideNum--
    },
    calculateSlideMaxSize () {
      const leftRightMargin = 100
      const topBottomMargin = 120
      this.slideMaxWidth = window.innerWidth - leftRightMargin
      this.slideMaxHeight = window.innerHeight - topBottomMargin
    }
  },
  computed: {
    slideMarkdown () {
      const markdown = this.markdown || this.$store.getters.markdown
      return this.$_getSlideMarkdown(markdown)
    }
  }
}
</script>

<style lang="scss">
  // html {
  //   background-color: map-get($color-brand, "main");
  // }

  // #app {
  //   min-width: 0;
  //   min-height: 0;
  // }
</style>

<style scoped lang="scss">
  $margin: 5%;

  .presentation {
    box-sizing: border-box;
    position: relative;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    overflow: hidden;
    width: 100vw;
    height: 100vh;
    background-color: map-get($color-brand, "main");

    // Layout
    // ==============================
    &-header {
      margin: 20px;
    }

    &-body {
      display: flex;
      justify-content: center;
      margin: 0 $margin;
    }

    &-footer {
      margin: 20px $margin;
      height: 20px;
    }

    // Components
    // ==============================
    &-close {
      $color: lighten( map-get($color-brand, "main-bright"), 2.5% );

      font-size: 1.0rem;
      color: $color;

      a {
        display: flex;
        align-items: center;
        overflow: hidden;
        color: inherit;
        text-decoration: none;
      }

      &-button {
        padding: 5px 10px;
        margin-right: 5px;
        font-family: $ff-mono;
        color: map-get($color-brand, "main");
        font-weight: bold;
        background-color: $color;
        border-radius: 3px;
      }
    }

    &-pager {
      $top: 3px;
      $left: 6px;

      display: flex;
      margin-top: -#{$top};
      margin-left: -#{$left};

      >li {
        margin-top: $top;
        margin-left: $left;

        // &:first-child {
          // margin-left: 0;
        // }
      }

      a {
        display: block;
        overflow: hidden;
        width: 10px;
        height: 6px;
        text-indent: 100%;
        white-space: nowrap;
        background-color: #fff;
        opacity: 0.2;
        transition-duration: $duration;

        &.is-now {
          opacity: 1;
        }

        &:not(.is-now) {
          &:hover {
            opacity: 0.5;
          }
        }
      }
    }

    &-logo {
      position: absolute;
      right: 10px;
      bottom: 10px;
      width: 100px;

      a {
        display: block;
        overflow: hidden;
        transition-duration: $duration;

        &:hover {
          opacity: $alpha;
        }
      }
    }
  }
</style>
