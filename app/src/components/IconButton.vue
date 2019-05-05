<template>
  <div>
    <router-link v-if="to && Object.keys(to).length" class="icon-button" :class="[balloonPositionClass, iconColorClass, iconReverseClass]" :to="to">
      <component :is="icon" />
      <span class="icon-button-balloon">
        {{balloonText}}
      </span>
    </router-link>

    <router-link v-else class="icon-button" :class="[balloonPositionClass, iconColorClass, iconReverseClass]" sto="javascript:void(0)">
      <component :is="icon" />
      <span class="icon-button-balloon">
        {{balloonText}}
      </span>
    </router-link>
  </div>
</template>

<script>
export default {
  name: 'IconButton',
  props: {
    svg: String,
    to: {
      type: [String, Object],
      default: () => {}
    },
    balloonPosition: String,
    balloonText: String,
    iconColor: String,
    iconReverse: Boolean
  },
  computed: {
    balloonPositionClass () {
      if (this.balloonPosition === 'left') {
        return ['is-left']
      } else if (this.balloonPosition === 'right') {
        return ['is-right']
      } else {
        return []
      }
    },
    iconColorClass () {
      if (this.iconColor) {
        return ['is-'+this.iconColor]
      } else {
        return []
      }
    },
    iconReverseClass () {
      if (this.iconReverse === true) {
        return ['is-reverse']
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
    cursor: pointer;

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
      svg,
      svg path {
        fill: map-get($color-brand, "gray-light");
      }
    }

    &.is-reverse {
      svg {
        transform: scale(-1);
        transform-origin: center center;
      }
    }

    &.is-vertical-reverse {
      svg {
        transform: scale(1, -1);
      }
    }

    &.is-horizontal-reverse {
      svg {
        transform: scale(-1, 1);
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
