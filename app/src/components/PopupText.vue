<template>
  <div class="popup-text">
    <slot />
  </div>
</template>

<script>
export default {
  name: 'PopupText',
  props: {
  }
}
</script>

<style scoped lang="scss">
  .popup-text {
    position: relative;
    top: 10px;
    box-sizing: border-box;
    display: flex;
    justify-content: center;
    align-items: center;
    width: auto;
    height: 30px;
    padding: 5px 15px;
    border: 1px solid map-get($color-brand, "accent");
    border-radius: 15px;
    opacity: 0;
    pointer-events: none;
    transition-duration: $duration * 2;

    a {
      color: inherit;

      &:hover {
        text-decoration: none;
      }
    }

    &.is-active {
      top: 0;
      opacity: 1;
      pointer-events: auto;
      transition-duration: $duration;
    }
  }
</style>
