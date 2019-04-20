<template>
  <div>
    <router-link v-if="to && Object.keys(to).length" class="text-button" :class="buttonStylesClass" :to="to" target="_blank">
      <component :is="icon" />
      {{text}}
    </router-link>

    <div v-else class="text-button" :class="buttonStylesClass">
      <component :is="icon" />
      {{text}}
    </div>
  </div>
</template>

<script>
export default {
  name: 'TextButton',
  props: {
    styleClass: String,
    svg: String,
    to: Object,
    text: String
  },
  computed: {
    buttonStylesClass () {
      if (this.styleClass) {
        return ['is-'+this.styleClass]
      } else {
        return []
      }
    },
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
    cursor: pointer;

    svg {
      margin-right: 5px;

      path {
        fill: map-get($color-brand, "main");
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
