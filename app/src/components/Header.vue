<template>
  <header>
    <div class="l-header">
      <h1 class="header-logo">
        <BrandLogo />
      </h1>
      <div class="header-tools">
        <div class="header-inner">
          <UserSlideListShowButton />
        </div>
        <div class="header-inner">
          <PresentationButton />
          <UserIcon />
          <LoginButton />
          <LogoutButton />
        </div>
      </div>
    </div>
  </header>
</template>

<script>
import BrandLogo from '@/components/BrandLogo.vue'
import UserSlideListShowButton from '@/components/UserSlideListShowButton.vue'
import UserIcon from '@/components/UserIcon.vue'
import LoginButton from '@/components/LoginButton.vue'
import LogoutButton from '@/components/LogoutButton.vue'

export default {
  name: 'Header',
  components: {
    BrandLogo,
    UserSlideListShowButton,
    UserIcon,
    LoginButton,
    LogoutButton
  },
  props: {
  },
  data () {
    return {
    }
  },
  computed: {
  }
}
</script>

<style scoped lang="scss">
.l-header {
  position: relative;
  height: 50px;
  background-color: #fff;
}

.header {
  &-tools {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  &-inner {
  }
}
</style>
