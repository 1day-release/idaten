<template>
  <div>
    <Slide class="slide" v-for="(value, index) in slideMarkdown" :key="index" :markdown="value" />
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
.slide {
  border: 1px solid gray;
}
</style>
