<template>
  <div class="preview">
    <ul class="preview-list">
      <li v-for="(value, index) in slideMarkdown" :key="index">
        <Slide :markdown="value" />
      </li>
    </ul>
  </div>
</template>

<script>
import Slide from '@/components/Slide.vue'

export default {
  name: 'Preview',
  components: {
    Slide
  },
  props: {
  },
  data () {
    return {
    }
  },
  computed: {
    slideMarkdown () {
      return this.$_getSlideMarkdown(this.$store.getters.markdown)
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
      $position: 20px;
      $padding: 40px;

      position: absolute;
      top: 0;
      bottom: 0;
      left: 0;
      overflow-y: auto;
      min-width: calc(100% + 17px);
      width: calc(100% - #{$position * 3});
      max-width: $pc-min-width * 0.35 - $position;
      height: auto;
      max-height: 100%;
      padding-right: $position;
      margin: auto;

      >li {
        width: 100%;
        padding-bottom: $padding;

        &:first-child {
          padding-top: $padding;
        }
      }
    }
  }
</style>
