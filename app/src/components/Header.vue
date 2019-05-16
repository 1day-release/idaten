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
          <li>
            <div class="button-block">
              <div class="button-mask is-logout js-popup-trigger js-popup-timeout">
                <ButtonMask />
              </div>
              <IconButton
                balloon-position="left"
                balloon-text="スライドリスト"
                svg="@/assets/icon-slidelist.svg"
                class="user-slide-list-button"
                @click.native="showUserSlideList"
              />
            </div>
          </li>
          <li>
            <div class="save-status">
              <p class="save-status-text">
                <span>更新しました</span>
                <span class="is-saving">更新中…</span>
              </p>
              <UpdateIcon class="save-status-icon" />
            </div>
          </li>
        </ul>
        <ul class="header-tool">
          <li>
            <IconButton
              balloon-text="プレゼンモード"
              svg="@/assets/icon-presentation.svg"
              :to="`/presentation/${activeSlideId}/1`"
              class="presentation-button"
            />
          </li>
          <li>
            <div class="button-block">
              <div class="button-mask is-logout js-popup-trigger js-popup-timeout">
                <ButtonMask />
              </div>
              <IconButton
                balloon-text="スライドをシェア"
                svg="@/assets/icon-link.svg"
              />
            </div>
            <PopupSetting
              id="popup-setting-share"
              style="right:-191px;"
            >
              <aside>
                <dl class="setting-item">
                  <dt>ボードの共有設定</dt>
                  <dd>
                    <input
                      type="radio"
                      class="input-radio"
                      name="board-share"
                      id="board-share-private"
                      checked
                    >
                    <label for="board-share-private">
                      <span>
                        <strong class="color-text-sub">自分だけが</strong>アクセスできる
                      </span>
                    </label>
                  </dd>
                  <dd>
                    <input
                      type="radio"
                      class="input-radio"
                      name="board-share"
                      id="board-share-view"
                    >
                    <label for="board-share-view">
                      <span>
                        リンクを知ってる全員が<strong class="color-text-sub">閲覧できる</strong>
                      </span>
                    </label>
                  </dd>
                  <dd>
                    <input
                      type="radio"
                      class="input-radio"
                      name="board-share"
                      id="board-share-edit"
                    >
                    <label for="board-share-edit">
                      <span>
                        リンクを知ってる全員が<strong class="color-text-sub">編集できる</strong>
                      </span>
                    </label>
                  </dd>
                </dl>
                <hr>
                <input
                  type="text"
                  class="input-text"
                  name=""
                  id=""
                  value="ここに今いるURLぶち込んでおくれーここに今いるURLぶち込んでおくれーここに今いるURLぶち込んでおくれー"
                >
                <ul class="button-list">
                  <li>
                    <a
                      class="submit-button"
                      href="javascript:void(0)"
                    >
                      リンクをコピー
                    </a>
                  </li>
                </ul>
              </aside>
            </PopupSetting>
          </li>
          <li>
            <IconButton
              balloon-text="設定"
              svg="@/assets/icon-setting.svg"
            />
            <PopupSetting
              id="popup-setting-slide"
              style="right:-124px;"
            >
              <aside>
                <dl class="setting-item">
                  <dt>ボードの削除</dt>
                  <dd>
                    <ul class="button-list">
                      <li>
                        <a
                          class="submit-button is-red"
                          href="javascript:void(0)"
                        >
                          このボードを削除する
                        </a>
                      </li>
                      <li>
                        <a
                          class="submit-button is-link"
                          href="javascript:void(0)"
                        >
                          キャンセル
                        </a>
                      </li>
                    </ul>
                  </dd>
                </dl>
              </aside>
            </PopupSetting>
          </li>
          <li class="is-separate is-login">
            <IconButton
              balloon-text="ログアウト"
              svg="@/assets/icon-logout.svg"
              class="logout-button"
            />
          </li>
          <li class="is-login">
            <UserIcon
              balloon-position="right"
              user-name="ワンデイ太郎"
              email="1day-release@gmail.com"
            />
          </li>
          <li class="is-separate is-logout">
            <TextButton
              text="ログイン"
              svg=""
              class="login-button"
              @click.native="login"
            />
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
import UpdateIcon from '@/assets/icon-refresh.svg'
import ButtonMask from '@/assets/icon-mask.svg'
import PopupSetting from '@/components/PopupSetting.vue'

export default {
  name: 'Header',
  components: {
    BrandLogo,
    IconButton,
    TextButton,
    UserIcon,
    UpdateIcon,
    ButtonMask,
    PopupSetting
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
    showUserSlideList () {
      this.$store.commit('slides')
      this.$store.dispatch('changeStateUserSlideList', !this.$store.getters.userSlideListState)
    },
    login () {
      const redirectURI = location.origin
      location.href = `https://accounts.google.com/o/oauth2/auth?client_id=87857787925-1ojg7q93r5fabr94uqfs9e0165tmrh1m.apps.googleusercontent.com&scope=openid email profile &response_type=code&redirect_uri=${redirectURI}&state=idaten&access_type=offline&approval_prompt=force`
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
          // &:last-child {
            // padding-right: $padding + 1px;
            // padding-left: $padding;

            // &::after {
            //   content: "";
            //   right: 0;
            // }
          // }
        }
      }
    }
  }

  .save-status {
    display: flex;
    justify-content: space-between;
    align-items: center;

    &-text {
      margin-right: 5px;
      font-size: 1.0rem;
      color: map-get($color-brand, "text-main");

      >span {

        &.is-saving {
          display: none;
        }
      }
    }

    &-icon {
      transform-origin: 58% 50%;
    }

    &.is-saving {
      .save-status {
        &-text {
          >span {

            &:not([class]) {
              display: none;
            }

            &.is-saving {
              display: block;
            }
          }
        }

        &-icon {
          animation: spinLeft 1.4s linear infinite;
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
