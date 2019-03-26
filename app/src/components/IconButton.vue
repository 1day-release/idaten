<template>
  <router-link class="presentation-button" :to="to" target="_blank">
    <component :is="icon" />
    <span class="presentation-button-text">
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
  .presentation-button {
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
    width: 36px;
    height: 36px;
    color: inherit;
    text-decoration: none;

    svg {
      path {
        transition-duration: $duration;
        // fill: map-get($color-brand, "sub");
      }
    }

    &-text {
      box-sizing: border-box;
      position: absolute;
      top: calc(100% + #{15px - 1px - 5px});
      right: -165%;
      left: -165%;
      display: flex;
      justify-content: center;
      align-items: center;
      height: 26px;
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
        right: 0;
        left: 0;
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

    &:hover {
      svg {
        path {
          fill: map-get($color-brand, "sub");
        }
      }

      .presentation-button {
        &-text {
          top: calc(100% + #{15px - 1px});
          opacity: 1;
        }
      }
    }
  }
</style>
