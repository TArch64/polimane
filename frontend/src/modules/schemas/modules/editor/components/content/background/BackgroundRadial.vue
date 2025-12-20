<template>
  <pattern
    x="0"
    y="0"
    patternUnits="userSpaceOnUse"
    :width="BEAD_SIZE * 2"
    :height="BEAD_SIZE * 2"
  >
    <rect
      v-for="(pos, index) in patternPositionsFront"
      :key="index"
      :x="pos.x"
      :y="pos.y"
      :width="BEAD_SIZE"
      :height="BEAD_SIZE"
      :fill="fillColor"
      fill-opacity="0.2"
      :stroke="strokeColor"
      stroke-width="1"
      class="canvas-content__background-bead"
    />
  </pattern>
</template>

<script setup lang="ts">
import { BEAD_SIZE } from '@editor/const';
import { computed } from 'vue';
import { useEditorStore } from '@editor/stores';

const editorStore = useEditorStore();

const shift = 0;

const patternPositionsFront = [
  { x: shift, y: shift },
  { x: BEAD_SIZE + shift, y: shift },
  { x: (BEAD_SIZE / 2) + shift, y: BEAD_SIZE + shift },
];

const emptyColor = '#CFCFCF';
const backgroundColor = computed(() => editorStore.schema.backgroundColor);

const fillColor = computed(() => {
  return `color-mix(in srgb, ${backgroundColor.value}, ${emptyColor})`;
});

const strokeColor = computed(() => {
  return `color-mix(in srgb, ${backgroundColor.value}, ${emptyColor} 50%)`;
});
</script>

<style scoped>
@layer page {
  .canvas-content__background-bead {
    transition: fill 0.15s ease-out;
  }
}
</style>
