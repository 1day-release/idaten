<template>
  <div class="user" :class="balloonClass">
    <span class="user-icon">
      <img src="//placehold.jp/150x150.png" alt="">
    </span>
    <span class="user-balloon">{{userName}}&nbsp;&lt;{{email}}&gt;<!-- name --></span>
  </div>
</template>

<script>
export default {
  name: 'UserIcon',
  props: {
    balloonPosition: String,
    userName: String,
    email: String
  },
  computed: {
    balloonClass () {
      if (this.balloonPosition === 'left') {
        return ['is-left']
      } else if (this.balloonPosition === 'right') {
        return ['is-right']
      } else {
        return []
      }
    }
  }
}
</script>

<style scoped lang="scss">
  .user {
    $size: 30px;

    position: relative;
    display: flex;
    justify-content: center;
    width: $size;
    height: $size;
    padding: 3px;
    // cursor: pointer;

    &-icon {
      position: relative;
      display: flex;
      justify-content: center;
      align-items: center;
      overflow: hidden;
      width: 100%;
      border-radius: 50%;

      &::before,
      &::after {
        content: "";
        display: block;
      }

      &::before {
        padding-top: 100%;
      }

      &::after {
        box-sizing: border-box;
        position: absolute;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
        margin: auto;
        border: 2px solid map-get($color-brand, "sub");
        border-radius: inherit;
        pointer-events: none;
        opacity: 0;
        transition-duration: $duration;
      }

      >img {
        position: absolute;
        top: 0;
        left: 0;
        display: block;
        width: 100%;
        height: auto;
      }
    }

    &-balloon {
      @include balloon($size, 3px);
    }

    @include balloon-styles(3px);

    &:hover {

      .user {
        &-icon {

          &::after {
            opacity: 1;
          }
        }

        &-balloon {
          @include balloon-hover-styles;
        }
      }
    }
  }
</style>
