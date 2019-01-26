<template>
  <div class="preview">
    <ul class="preview-list">
      <li>
        <Slide />
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
  }
}
</script>

<style scoped lang="scss">
  .preview {
    position: relative;
    width: 100%;
    height: 100%;

    &-list {
      position: absolute;
      left: 20px;

      >li {
      }
    }
  }
</style>
