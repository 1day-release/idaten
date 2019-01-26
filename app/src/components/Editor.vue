<template>
  <div class="l-editor">
    <div class="editor">
      <code></code>
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
  }
}
</script>

<style scoped lang="scss">
.l-editor {}

.editor {

  code {
  }
}
</style>
