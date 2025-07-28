<template>
  <span
    v-for="index in dots"
    class="button-loading__dot"
    :key="index"
    :style="{ '--button-dot-index': index - 1 }"
  />
</template>

<script setup lang="ts">
const dots = 3;
</script>

<style scoped>
@layer components {
  .button-loading__dot {
    position: absolute;
    width: var(--button-dot-size);
    height: var(--button-dot-size);
    border-radius: 100%;
    background-color: currentColor;
    left: calc((50% - (var(--button-dots-width) / 2)) + (var(--button-dot-index) * (var(--button-dot-gap) + var(--button-dot-size))));
    animation: button-loading-dot 1s calc(1s * var(--button-dot-index)) linear infinite;
    scale: calc(1 - var(--button-dot-index) * (0.5 / v-bind("dots")));
    --button-dot-size: 8px;
    --button-dot-gap: 4px;
    --button-dots-width: calc(v-bind("dots") * var(--button-dot-size) + (v-bind("dots") - 1) * var(--button-dot-gap));
  }

  @keyframes button-loading-dot {
    from {
      scale: calc(1 - var(--button-dot-index) * (0.5 / v-bind("dots")))
    }
    50% {
      scale: calc(0.5 + var(--button-dot-index) * (0.5 / v-bind("dots")))
    }
  }
}
</style>
