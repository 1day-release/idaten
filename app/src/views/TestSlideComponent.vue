<template>
  <div>
    <div v-for="(testCase, index) in slideTestCases" :key="index">
      <Slide :html="testCase.html" />
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
      console.log(data)
      this.slideTestCases = data
    })
  }
}
</script>

<style scoped lang="scss">
</style>
