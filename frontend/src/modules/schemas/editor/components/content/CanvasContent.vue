<template>
  <g :transform class="canvas-content" v-on="listeners">
    <rect
      opacity="0"
      :x="gridSize.minX"
      :y="gridSize.minY"
      :width="gridSize.width"
      :height="gridSize.height"
    />

    <CanvasSector
      v-for="{ sector, grid } of sectors"
      :key="sector"
      :grid="grid.value"
    />
  </g>
</template>

<script setup lang="ts">
import { useBeadPainting, useBeadsGrid } from '../../composables';
import { useEditorStore } from '../../stores';
import CanvasSector from './CanvasSector.vue';

const props = defineProps<{
  wrapperRect: DOMRect;
}>();

const editorStore = useEditorStore();

const listeners = useBeadPainting();
const { sectors, gridSize } = useBeadsGrid(() => editorStore.schema);

const transform = (() => {
  const y = (props.wrapperRect.height - gridSize.height) / 2;
  const x = (props.wrapperRect.width - gridSize.width) / 2;

  return `translate(${x}, ${y})`;
})();
</script>

<style scoped>
@layer page {
  .canvas-content:hover {
    cursor: crosshair;
  }
}
</style>
