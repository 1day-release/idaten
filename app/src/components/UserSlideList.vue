<template>

  <div class="user-slide">
    <div class="user-slide-head">
      <div class="user-slide-head-item" v-on:click="showUserSlideList()">
        <IconButton balloon-position="left" icon-color="light-gray" icon-reverse balloon-text="スライドリストを閉じる" svg="@/assets/icon-slidelist.svg"/>
      </div>
      <div class="user-slide-head-item">
        <TextButton class="is-ghost" text="新規スライド作成" svg="@/assets/icon-update.svg" :to="{name: 'createSlide'}" />
      </div>
    </div>
    <ul class="user-slide-list">
      <li>
        <a class="slide-cover" href="#">
          <Slide :max-width="slideMaxWidth" :max-height="slideMaxHeight" />
        </a>
      </li>
      <li class="is_now">
        <a class="slide-cover" href="#">
          <Slide />
        </a>
      </li>
      <li>
        <a class="slide-cover" href="#">
          <Slide />
        </a>
      </li>
      <li>
        <a class="slide-cover" href="#">
          <Slide />
        </a>
      </li>
      <li>
        <a class="slide-cover" href="#">
          <Slide />
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
import IconButton from '@/components/IconButton.vue'
import TextButton from '@/components/TextButton.vue'
import Slide from '@/components/Slide.vue'

export default {
  name: 'UserSlideList',
  components: {
    IconButton,
    TextButton,
    Slide
  },
  props: {
  },
  data () {
    return {
    }
  },
  computed: {
    isActive () {
      console.log(this.$store.getters.userSlideListState)
      return this.$store.getters.userSlideListState
    }
  },
  methods: {
    showUserSlideList: function (e) {
      console.log('test')
      this.$store.dispatch('changeStateUserSlideList', !this.$store.getters.userSlideListState)
    }
  }
}
</script>

<style scoped lang="scss">

  $padding: 20px;
  $border-color: map-get($color-brand, "main-bright");

  .user-slide {
    display: block;
    overflow: hidden;
    width: 100%;
    height: 100%;
    background-color: map-get($color-brand, "main-light");

    &-head {
      position: relative;
      display: flex;
      justify-content: center;
      align-items: center;
      padding-bottom: 16px;
      margin: 15px 0 10px;

      &::after {
        content: "";
        position: absolute;
        right: $padding;
        bottom: 0;
        left: $padding;
        height: 1px;
        margin: 0 auto;
        background-color: $border-color;
      }

      &-item {
        position: relative;
        padding-left: 11px;
        margin-left: 10px;

        &::before {
          content: "";
          position: absolute;
          top: 0;
          bottom: 0;
          left: 0;
          width: 1px;
          height: 30px;
          background-color: $border-color;
        }

        &:first-child {
          padding-left: 0;
          margin-left: 0;

          &::before {
            content: none;
          }
        }
      }
    }

    &-list {
      &-hide {

        &-button {
          appearance: none;
          display: flex;
          justify-content: center;
          align-items: center;
          width: 36px;
          height: 36px;
          padding: 0;
          border: none;
          background-color: transparent;
          cursor: pointer;

          svg {
            transform: scale(-1, 1);

            path {
              transition-duration: $duration;
            }
          }

          &:hover {

            svg {
              path {
                fill: map-get($color-brand, "sub");
              }
            }
          }
        }
      }

      box-sizing: border-box;
      display: block;
      overflow-y: scroll;
      width: auto;
      min-width: calc(100% + 17px);
      height: 100%;
      padding-right: 17px;

      >li {
        box-sizing: border-box;
        width: inherit;
        padding: 10px 20px;

        &:not(.is_now) {
          transition-duration: $duration;

          &:hover {
            background-color: rgba( map-get($color-brand, "main-bright"), 0.5 );
          }
        }

        &.is_now {
          position: relative;
          background-color: map-get($color-brand, "main-bright");

          &::before {
            $size: 6px;

            content: "";
            position: absolute;
            top: 0;
            bottom: 0;
            left: 7px;
            display: block;
            width: $size;
            height: $size;
            margin: auto 0;
            background-color: map-get($color-brand, "sub");
            border-radius: 50%;
          }
        }
      }
    }
  }

  .slide-cover {
    display: block;
    overflow: hidden;
  }
</style>
