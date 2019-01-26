<template>
  <div>
    <div class="editor">
      <textarea v-model="markdown" @input="editorInput"></textarea>
    </div>
    <Popups>
      <PopupLogout />
    </Popups>
  </div>
</template>

<script>
import Popups from '@/components/Popups.vue'
import PopupLogout from '@/components/PopupLogout.vue'

export default {
  name: 'Editor',
  components: {
    Popups,
    PopupLogout
  },
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
  .editor {

    code {
    }
  }
</style>
