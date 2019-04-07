<template>
  <router-link class="text-button" :to="to" target="_blank">
    <component :is="icon" />
    {{text}}
  </router-link>
</template>

<script>
export default {
  name: 'TextButton',
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
  .text-button {
    $height: 30px;
    $color: map-get($color-brand, "accent");

    box-sizing: border-box;
    display: flex;
    justify-content: center;
    align-items: center;
    width: auto;
    min-width: 95px;
    height: $height;
    padding: 0 20px;
    border: 1px solid $color;
    font-size: 1.0rem;
    color: inherit;
    text-decoration: none;
    background-color: $color;
    border-radius: $height / 2;
    transition-duration: $duration;

    svg {
      margin-right: 5px;

      path {
        // fill: $color;
        transition-duration: $duration;
      }
    }

    &:not(.is-ghost) {
      &:hover {
        opacity: $alpha;
      }
    }

    &:hover {

      svg {
        path {
          // fill: map-get($color-brand, "sub");
        }
      }
    }

    &.is-ghost {
      color: $color;
      background-color: transparent;

      svg {
        path {
          margin-right: 5px;
          fill: $color;
        }
      }

      &:hover {
        color: map-get($color-brand, "text");
        background-color: $color;
        // background-color: rgba($color, 0.15);

        svg {
          path {
            fill: map-get($color-brand, "text");
          }
        }
      }
    }
  }
</style>
