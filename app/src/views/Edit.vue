<template>
  <div>
    <Header />
    <Main>
      <UserSlideList />
      <Editor />
      <Preview />
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
  name: 'home',
  components: {
    Header,
    Main,
    UserSlideList,
    Editor,
    Preview
  }
}
</script>
