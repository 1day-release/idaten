<template>
  <div>
    <div
      class="test-block"
      v-for="(testCase, index) in slideTestCases"
      :key="index"
    >
      <Slide
        :markdown="testCase.markdown"
        :page-number="testCase.pageNumber"
        :width="1112"
      />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
import Slide from '@/components/Slide.vue'

export default {
  name: 'Presentation',
  components: {
    Slide
  },
  data () {
    return {
      slideTestCases: []
    }
  },
  created () {
    fetch('https://script.google.com/macros/s/AKfycbxYk9NMXJbjxCdHMyTmoFH-mIVfoCgw2PXr_0NWIuFvp5dzfrE/exec').then(response => response.json()).then((data) => {
      this.slideTestCases = data
    })
  }
}
</script>

<style scoped lang="scss">
  .test-block {
    background-color: map-get($color-brand, "main");
  }
</style>
