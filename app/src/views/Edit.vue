<template>
  <div>
    <Header />
    <Main>
      <UserSlideList />
      <Editor class="editor" />
      <Preview class="preview" />
    </Main>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
import Header from '@/components/Header.vue'
import Main from '@/components/Main.vue'
import UserSlideList from '@/components/UserSlideList.vue'
import Editor from '@/components/Editor.vue'
import Preview from '@/components/Preview.vue'

export default {
  name: 'Edit',
  components: {
    Header,
    Main,
    UserSlideList,
    Editor,
    Preview
  }
}
</script>

<style scoped lang="scss">
.editor {
  float: left;
  width: 60vw;
  height: 70vh;
  overflow: scroll;
}
.preview {
  float: left;
  width: 40vw;
  height: 70vh;
  overflow: scroll;
}
</style>
