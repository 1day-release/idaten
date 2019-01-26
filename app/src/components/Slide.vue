<template>
  <ul class="slide">
    <li>
      <SlidePage />
    </li>
  </ul>
</template>

<script>
import SlidePage from '@/components/SlidePage.vue'

export default {
  name: 'Slide',
  components: {
    SlidePage
  },
  props: {
  }
}
</script>

<style scoped lang="scss">
</style>
