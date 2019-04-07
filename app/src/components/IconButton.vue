<template>
  <div>
    <router-link v-if="to" class="presentation-button" :to="to" target="_blank">
      <component :is="icon" />
      <span class="presentation-button-text">
        {{text}}
      </span>
    </router-link>

    <div v-else>
      <component :is="icon" />
      <span class="presentation-button-text">
        {{text}}
      </span>
    </div>
  </div>
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

    &-balloon {
      @include balloon($size);
    }

    @include balloon-styles();

    &.is-light-gray {
      svg {
        fill: map-get($color-brand, "gray-light");
      }
    }

    &.is-reverse {
      svg {
        transform: scale(-1, 1);
        transform-origin: center center;
      }
    }

    &:hover {
      svg {
        path {
          fill: map-get($color-brand, "sub");
        }
      }

      .icon-button {
        &-balloon {
          @include balloon-hover-styles();
        }
      }
    }
  }
</style>
