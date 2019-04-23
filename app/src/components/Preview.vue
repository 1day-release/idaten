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

    box-sizing: border-box;
    position: relative;
    overflow-x: hidden;
    display: block;
    width: calc(100% - #{$position * 2});
    max-width: $pc-min-width * 0.35 - $position;
    height: 100%;
    margin-left: $position;

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
      padding-right: 17px;
      margin: auto;

      >li {
        width: inherit;
        padding-bottom: $padding;

        &:first-child {
          padding-top: $padding;
        }
      }
    }
  }
</style>
