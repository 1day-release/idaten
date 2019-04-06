<template>
  <router-link class="icon-button" :to="to" target="_blank">
    <component :is="icon" />
    <span class="icon-button-text">
      {{text}}
    </span>
  </router-link>
</template>

<script>
export default {
  name: 'IconButton',
  props: {
    svg: String,
    to: Object,
    text: String
  },
  computed: {
    icon () {
      return () => import(
        /* webpackChunkName: "assets" */
        /* webpackMode: "lazy" */
        `@/${this.svg.replace('@/', '')}`
      )
    }
  }
}
</script>

<style scoped lang="scss">
  .icon-button {
    $size: 36px;
    $margin: 10px;

    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
    width: $size;
    height: $size;
    color: inherit;
    text-decoration: none;

    svg {
      path {
        transition-duration: $duration;
      }
    }

    &-text {
      box-sizing: border-box;
      position: absolute;
      top: calc(100% + #{15px - 1px - 5px});
      display: flex;
      justify-content: center;
      align-items: center;
      height: 26px;
      padding: 0 20px;
      margin: 0 auto;
      color: map-get($color-brand, "text-white");
      background-color: map-get($color-brand, "gray");
      border-radius: 13px;
      box-shadow: 4px 4px 11px 0 map-get($color-brand, "shadow");
      white-space: nowrap;
      pointer-events: none;
      opacity: 0;
      transition-duration: $duration;

      &::before {
        content: "";
        position: absolute;
        top: -5px;
        right: ($size + $margin) / 2;
        left: ($size + $margin) / 2;
        display: block;
        width: 0;
        height: 0;
        border-width: 0 3px;
        border-style: solid;
        border-color: transparent;
        border-bottom: 5px solid map-get($color-brand, "gray");
        margin: 0 auto;
      }
    }

    &.is-left & {
      &-text {
        left: -#{$margin};

        &::before {
          right: auto;
        }
      }
    }

    &.is-right & {
      &-text {
        right: -#{$margin};

        &::before {
          left: auto;
        }
      }
    }

    &:hover {
      svg {
        path {
          fill: map-get($color-brand, "sub");
        }
      }
    }

    &:hover & {
      &-text {
        top: calc(100% + #{15px - 1px});
        opacity: 1;
      }
    }
  }
</style>
