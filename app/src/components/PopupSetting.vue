<template>
  <div class="popup-setting">
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: 'PopupSetting',
  props: {
  }
}
</script>

<style scoped lang="scss">
  .popup-setting {
    $bgColor: #fff;

    box-sizing: border-box;
    position: absolute;
    z-index: 10;
    top: calc(100% + 10px);
    width: 370px;
    padding: 20px;
    background-color: $bgColor;
    box-shadow: 4px 0 11px 0 rgba(0, 0, 0, 0.3);

    opacity: 0;
    pointer-events: none;

    &::before,
    &::after {
      content: "";
      position: absolute;
      background-color: $bgColor;
      display: block;
    }

    &::before {
      bottom: calc(100% - 4px);
      width: 6px;
      height: 6px;
      box-shadow: inherit;
      transform: rotate(45deg);
      transform-origin: center center;
    }

    &::after {
      top: 0;
      width: 20px;
      height: 10px;
    }

    &.is-active {
      pointer-events: auto;
      opacity: 1;
    }

    .setting-item {

      >dt {
        margin: 10px 0;
        font-size: 1.0rem;
        color: #5a5a5a;
      }

      >dd {
        display: block;
        overflow: hidden;
        margin-top: 1px;
      }

      strong {
        font-weight: bold;
      }
    }

    hr {
      display: block;
      width: 100%;
      height: 1px;
      border: 0;
      margin: 20px 0 30px;
      background-color: rgba(#000, 0.1);
    }

    .button-list {
      margin: 10px 0;

      >li {
        margin-top: 20px;

        &:first-child {
          margin-top: 0;
        }
      }
    }
  }
</style>
