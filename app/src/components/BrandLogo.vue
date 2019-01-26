<template>
  <span class="brand-logo">Logo</span>
</template>

<script>
export default {
  name: 'BrandLogo',
  props: {
  }
}
</script>

<style scoped lang="scss">
  .brand-logo {
    display: block;
    width: 100%;
    height: 100%;
    background-color: map-get($color-brand, "main");
  }
</style>
