<template>
  <div class="preview">
    <ul class="preview-list">
      <li v-for="n in pageCount" :key="n" :class="{'is-active': activePageNumber === n}">
        <Slide :markdown="markdown" :page-number="n" :width="347" />
      </li>
    </ul>
  </div>
</template>

<script>
import Slide from '@/components/Slide.vue'
import IdatenCore from 'idaten-core'
const core = new IdatenCore()

export default {
  name: 'Preview',
  components: {
    Slide
  },
  props: {
    activePageNumber: {
      type: Number,
      default: -1
    }
  },
  data () {
    return {
    }
  },
  mounted () {
  },
  methods: {
  },
  computed: {
    pageCount () {
      return core.countPage(this.markdown)
    },
    markdown () {
      return this.$store.getters.markdown
    }
  }
}
</script>

<style scoped lang="scss">
  .preview {
    $position: 20px;
    $padding: 40px;

    // box-sizing: border-box;
    position: relative;
    overflow-x: hidden;
    display: block;
    width: 100%;
    height: 100%;

    &-list {
      box-sizing: border-box;
      position: absolute;
      top: 1px;
      bottom: 0;
      left: 0;
      overflow-y: scroll;
      min-width: calc(100% + 17px);
      width: auto;
      height: auto;
      max-height: 100%;
      padding-right: $position + 17px;
      padding-left: $position;
      margin: auto;

      >li {
        position: relative;
        width: inherit;
        margin-bottom: $padding;
        box-shadow: 4px 4px 11px map-get($color-brand, "shadow");
        transition-duration: $duration;

        &::before,
        &::after {
          content: "";
          position: absolute;
          z-index: 10;
          top: 0;
          bottom: 0;

          pointer-events: none;
          opacity: 0;
          transition-duration: inherit;
        }

        &::before {
          right: calc(100% + 8px);
          display: block;
          width: 0;
          height: 0;
          border-style: solid;
          border-width: 6px 5px;
          border-left-width: 0;
          border-color: transparent;
          border-right-color: #fff;
          margin: auto 0;

        }

        &::after {
          right: 0;
          left: 0;
          border: 2px solid #d3d3d3;
          margin: auto;
        }

        &:first-child {
          margin-top: $padding;
        }

        &.is-active {

          &::before,
          &::after {
            opacity: 1;
          }
        }
      }
    }
  }
</style>
