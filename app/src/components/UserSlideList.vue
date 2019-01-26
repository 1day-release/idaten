<template>
  <div class="l-user-slide">
    <ul class="user-slide-list">
      <li>
        <a class="user-slide" href="#">
          <SlidePage />
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
import SlidePage from '@/components/SlidePage.vue'

export default {
  name: 'UserSlideList',
  components: {
    SlidePage
  },
  props: {
  }
}
</script>

<style scoped lang="scss">
.l-user-slide {
}

.user-slide {
  &-list {

    >li {
    }
  }

  color: inherit;
  text-decoration: none;
}
</style>
