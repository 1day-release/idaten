<template>
  <div class="user-slide">
    <div class="user-slide-head">
      <div
        class="user-slide-head-item"
        @click="hideUserSlideList()"
      >
        <IconButton
          balloon-position="left"
          icon-color="light-gray"
          icon-reverse
          balloon-text="スライドリストを閉じる"
          svg="@/assets/icon-slidelist.svg"
        />
      </div>
      <div
        class="user-slide-head-item"
        @click="createSlide"
      >
        <TextButton
          style-class="ghost"
          text="新規スライド作成"
          svg="@/assets/icon-update.svg"
        />
      </div>
    </div>
    <ul class="user-slide-list">
      <li
        v-for="(slide, index) in slides"
        :key="index"
        :class="{'is-now': slide.id === activeSlideId}"
      >
        <a
          class="slide-delete-button"
          href="javascript:void(0)a"
          @click="deleteSlide(slide.id, slide.id === activeSlideId)"
        >スライドを削除する</a>
        <a
          class="slide-cover"
          href="javascript:void(0)"
          @click="selectSlide(slide.id)"
        >
          <Slide
            :markdown="slide.markdown"
            :width="180"
          />
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
      return this.$store.getters.userSlideListState
    },
    slides () {
      return this.$store.getters.slides
    },
    activeSlideId () {
      return this.$store.getters.slideId
    }
  },
  methods: {
    hideUserSlideList () {
      this.$store.commit('userSlideListState', false)
    },
    selectSlide (slideId) {
      this.$store.dispatch('setSlideId', slideId)
      this.hideUserSlideList()
    },
    async createSlide () {
      const slideId = await this.$store.dispatch('createSlide')
      this.$store.dispatch('setSlideId', slideId)
      this.hideUserSlideList()
    },
    async deleteSlide (slideId, ownFlg) {
      if (!window.confirm('選択したスライドを削除しますか?')) return
      await this.$store.dispatch('deleteSlide', slideId)
      await this.$store.dispatch('reloadList')
      if (ownFlg) {
        const slideId = this.$store.getters.slides[0].id
        this.$store.dispatch('setSlideId', slideId)
        this.$router.push('/' + slideId)
      }
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
        position: relative;
        width: inherit;
        padding: 10px 20px;

        &:not(.is-now) {
          transition-duration: $duration;

          &:hover {
            background-color: rgba( map-get($color-brand, "main-bright"), 0.5 );
          }
        }

        &:hover {

          .slide-delete-button {
            opacity: 1;
            // pointer-events: auto;
          }
        }

        &.is-now {
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

  .slide {
    &-delete-button {
      $size: 20px;
      $borderWidth: 2px;
      $borderColor: lighten(#f00, 30%);

      box-sizing: border-box;
      position: absolute;
      z-index: 10;
      top: 10px - $size/2;
      right: 20px - $size/2;
      display: block;
      overflow: hidden;
      width: $size;
      height: $size;
      border: $borderWidth solid $borderColor;
      text-indent: 100%;
      white-space: nowrap;
      background-color: red;
      // background-color: map-get($color-brand, "accent");
      border-radius: 50%;
      transition-duration: $duration;
      opacity: 0;
      // pointer-events: none;

      &::before,
      &::after {
        content: "";
        position: absolute;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
        width: $borderWidth;
        height: 50%;
        margin: auto;
        background-color: $borderColor;
        transform-origin: center center;
      }

      &::before {
        transform: rotate(-45deg);
      }

      &::after {
        transform: rotate(45deg);
      }
    }

    &-cover {
      display: block;
      overflow: hidden;
      text-decoration: none;
    }
  }
</style>
