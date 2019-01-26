<template>
  <div>
    <textarea v-model="markdown" @input="editorInput"></textarea>
  </div>
</template>

<script>
export default {
  name: 'Editor',
  props: {
  },
  data () {
    return {
      markdown: ''
    }
  },
  methods: {
    editorInput () {
      this.$store.commit('markdown', this.markdown)
      localStorage.setItem('idaten.markdown', this.markdown)
    }
  },
  created () {
    this.markdown = localStorage.getItem('idaten.markdown')
    this.editorInput()
  }
}
</script>

<style scoped lang="scss">
textarea {
  width: 100%;
  height: 100%;
}
</style>
