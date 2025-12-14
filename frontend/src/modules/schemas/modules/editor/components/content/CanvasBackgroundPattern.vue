<template>
  <pattern
    x="0"
    y="0"
    patternUnits="userSpaceOnUse"
    :width="BEAD_SIZE"
    :height="BEAD_SIZE"
  >
    <rect
      x="0"
      y="0"
      :width="BEAD_SIZE"
      :height="BEAD_SIZE"
      :fill="emptyColor"
      fill-opacity="0.2"
      :stroke="emptyColor"
      stroke-width="0.5"
      class="canvas-content__background-bead"
    />
  </pattern>
</template>

<script setup lang="ts">
import { useBackgroundCanvasColor } from '@editor/composables';
import { BEAD_SIZE } from '@editor/const';

const emptyColor = useBackgroundCanvasColor();
</script>

<style scoped>
@layer page {
  .canvas-content__background-bead {
    transition: fill 0.15s ease-out;
  }
}
</style>
