<template>
  <div class="l-container is-full">
    <div class="presentation">
      <div class="presentation-header">
        <p class="presentation-close">
          <router-link to="/">
            <span class="presentation-close-button">Esc</span>で編集ページに戻る
          </router-link>
        </p>
      </div>
      <div class="presentation-body" @click="nextPage" @contextmenu.prevent="prevPage">
        <Slide :markdown="markdown" :page-number="pageNumber" :max-width="pageMaxWidth" :max-height="pageMaxHeight" />
      </div>
      <div class="presentation-footer">
        <ul class="presentation-pager">
          <li v-for="n in pageCount" :key="n"><router-link :to="'/presentation/' + n" :class="{'is-now': n === pageNumber}">{{n}}</router-link></li>
        </ul>
        <p class="presentation-logo">
          <a href="/">
            <BrandLogo class="is-gray" />
          </a>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
import Slide from '@/components/Slide.vue'
import BrandLogo from '@/components/BrandLogo.vue'
import IdatenCore from 'idaten-core'
const core = new IdatenCore()

export default {
  name: 'Presentation',
  components: {
    Slide,
    BrandLogo
  },
  data () {
    return {
      pageNumber: 1,
      pageMaxWidth: 0,
      pageMaxHeight: 0
    }
  },
  created () {
    this.pageNumber = parseInt(this.$route.params.pageNumber)

    if (this.$route.query.mdUrl) {
      fetch(this.$route.query.mdUrl).then(request => request.text()).then((markdown) => {
        this.markdown = markdown
      })
    }

    const keyEvent = (event) => {
      if (event.key === 'ArrowRight' || event.key === 'Enter') {
        this.nextPage()
      } else if (event.key === 'ArrowLeft' || event.key === 'Backspace') {
        this.prevPage()
      } else if (event.key === 'Escape') {
        window.removeEventListener('keydown', keyEvent)
        this.$router.push({ path: '/' })
      }
    }
    window.addEventListener('keydown', keyEvent)

    window.addEventListener('resize', this.calculatePageMaxSize)
    this.calculatePageMaxSize()
  },
  methods: {
    nextPage () {
      if (this.pageNumber < this.pageCount) {
        this.$router.push('/presentation/' + (this.pageNumber + 1))
      }
    },
    prevPage () {
      if (this.pageNumber > 1) {
        this.$router.push('/presentation/' + (this.pageNumber - 1))
      }
    },
    calculatePageMaxSize () {
      const leftRightMargin = 100
      const topBottomMargin = 120
      this.pageMaxWidth = window.innerWidth - leftRightMargin
      this.pageMaxHeight = window.innerHeight - topBottomMargin
    }
  },
  computed: {
    pageCount () {
      return core.countPage(this.markdown)
    },
    markdown () {
      return this.$store.getters.markdown
    }
  },
  watch: {
    $route (to, from) {
      this.pageNumber = parseInt(this.$route.params.pageNumber)
    }
  }
}
</script>

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
