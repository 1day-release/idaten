<template>
  <span class="brand-logo">
    <logo />
  </span>
</template>

<script>
import Logo from '@/assets/logo.svg'

export default {
  name: 'BrandLogo',
  components: {
    Logo
  },
  props: {
  }
}
</script>

<style scoped lang="scss">
  .brand-logo {
    position: relative;
    display: block;
    width: 100%;

    &::before {
      content: "";
      display: block;
      padding-top: 80.878 / 424.057 * 100%;
    }

    svg {
      position: absolute;
      top: 0;
      right: 0;
      bottom: 0;
      left: 0;
      margin: auto;

      .color1 { fill: #281a41; }
      .color2 { fill: #00a193; }
      .color3 { fill: #00ad9c; }
      .color4 { fill: #ffce00; }
      .color5 { fill: #40c799; }
    }

    &.is-gray {
      svg {
        [class^=color] {
          fill: rgba(#fff, 0.15);
        }
      }
    }
  }
</style>
