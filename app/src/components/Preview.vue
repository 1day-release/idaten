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
    position: relative;
    overflow: hidden;
    display: block;
    width: 100%;
    height: 100%;

    &-list {
      $position: 20px;
      $padding: 40px;

      position: absolute;
      top: $position;
      bottom: $position;
      left: $position;
      overflow-y: auto;
      min-width: calc(100% + 17px);
      width: calc(100% - #{$position * 3});
      max-width: $pc-min-width * 0.35 - $position;
      height: auto;
      max-height: 100%;
      padding-right: $position;
      margin: auto;

      // @include scrollBarRight(7px, 7px, rgba(#fff, 0.3), rgba(#fff, 0.1));

      >li {
        width: 100%;
        margin-top: $position;

        // &:first-child {
        //   margin-top: 0;
        // }
      }
    }
  }
</style>
