<template>
  <header>
    <div class="l-header">
      <h1 class="header-logo">
        <a href="#">
          <BrandLogo />
        </a>
      </h1>
      <div class="header-tools">
        <ul class="header-tool">
          <li v-on:click="showUserSlideList()">
            <div class="button-block">
              <div class="button-mask is-logout js-popup-trigger js-popup-timeout">
                <ButtonMask />
              </div>

              <IconButton balloon-position="left" balloon-text="スライドリスト" svg="@/assets/icon-slidelist.svg" />
            </div>
          </li>
        </ul>
        <ul class="header-tool">
          <li>
            <IconButton balloon-text="プレゼンモード" svg="@/assets/icon-presentation.svg" :to="`/presentation/${activeSlideId}/1`" />
          </li>
          <li>
            <div class="button-block">
              <div class="button-mask is-logout js-popup-trigger js-popup-timeout">
                <ButtonMask />
              </div>
              <IconButton balloon-text="スライドをシェア" svg="@/assets/icon-link.svg" />
            </div>
          </li>
          <!-- li>
            <IconButton balloon-text="設定" svg="@/assets/icon-setting.svg" />
          </li -->
          <li class="is-separate is-login">
            <IconButton balloon-text="ログアウト" svg="@/assets/icon-logout.svg" />
          </li>
          <li class="is-login">
            <UserIcon balloon-position="right" user-name="ワンデイ太郎" email="1day-release@gmail.com" />
          </li>
          <li class="is-separate is-logout">
            <TextButton text="ログイン" svg="" />
          </li>
        </ul>
      </div>
    </div>
  </header>
</template>

<script>
import BrandLogo from '@/components/BrandLogo.vue'
import IconButton from '@/components/IconButton.vue'
import TextButton from '@/components/TextButton.vue'
import UserIcon from '@/components/UserIcon.vue'
import ButtonMask from '@/assets/icon-mask.svg'

export default {
  name: 'Header',
  components: {
    BrandLogo,
    IconButton,
    TextButton,
    UserIcon,
    ButtonMask
  },
  props: {
  },
  data () {
    return {
    }
  },
  computed: {
    activeSlideId () {
      return this.$store.getters.slideId
    }
  },
  methods: {
    showUserSlideList: function (e) {
      this.$store.commit('slides')
      this.$store.dispatch('changeStateUserSlideList', !this.$store.getters.userSlideListState)
    }
  }
}
</script>

<style scoped lang="scss">
  @at-root .l-container:not(.is-logout) .is-logout {
    display: none !important;
  }

  @at-root .l-container.is-logout .is-login {
    display: none !important;
  }

  .l-header {
    position: relative;
    height: $pc-header-height;
    background-color: map-get($color-brand, "base");
  }

  .header {
    &-logo {
      position: absolute;
      top: 0;
      right: 0;
      bottom: 0;
      left: 0;
      display: flex;
      justify-content: center;
      align-items: center;
      width: 100px;
      margin: auto;

      >a {
        display: block;
        overflow: hidden;
        width: inherit;
        transition-duration: $duration;

        &:hover {
          opacity: $alpha;
        }
      }
    }

    &-tools {
      display: flex;
      justify-content: space-between;
      align-items: center;
      height: 100%;
    }

    &-tool {
      $padding: 15px;

      display: flex;
      height: inherit;

      >li {
        box-sizing: border-box;
        flex: 0 0 auto;
        position: relative;
        display: flex;
        justify-content: center;
        align-items: center;
        width: auto;
        min-width: 64px;
        padding-right: $padding;
        padding-left: $padding + 1px;
        height: inherit;

        &::before,
        &::after {
          position: absolute;
          top: 0;
          bottom: 0;
          margin: auto 0;
          width: 1px;
          height: 30px;
          background-color: map-get($color-brand, "border-light");
        }

        &::before {
          content: "";
          left: 0;
        }

        &:first-child {

          &::before {
            content: none;
          }
        }

        &.is-separate {
          &::before,
          &::after {
            height: auto;
          }
        }
      }

      &:first-child {
        >li {
          &:last-child {
            padding-right: $padding + 1px;
            padding-left: $padding;

            &::after {
              content: "";
              right: 0;
            }
          }
        }
      }
    }
  }

  .button {
    &-block {
      position: relative;
    }

    &-mask {
      position: absolute;
      z-index: 10;
      top: 0;
      right: 0;
      bottom: 0;
      left: 0;
      display: flex;
      justify-content: center;
      align-items: center;
      margin: auto;
      cursor: not-allowed;

      svg {
        .cls-1, .cls-3 {
          fill: none;
        }

        .cls-2 {
          clip-path: url(#clip-path);
        }

        .cls-3 {
          stroke: #cbcdd6;
          stroke-width: 0.5px;
        }
      }
    }
  }

  .user {
  }
</style>
