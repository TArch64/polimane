<template>
  <g
    class="canvas-content"
    :transform="gTransform"
    @mousedown="paint"
    v-on="listeners"
  >
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
import { computed } from 'vue';
import type { SchemaBeadCoord } from '@/models';
import { useBeadsGrid } from '../../composables';
import { useBeadsStore, useEditorStore, usePaletteStore } from '../../stores';
import CanvasSector from './CanvasSector.vue';

const props = defineProps<{
  wrapperRect: DOMRect;
}>();

const editorStore = useEditorStore();
const paletteStore = usePaletteStore();
const beadsStore = useBeadsStore();

const { sectors, gridSize } = useBeadsGrid(() => editorStore.schema);

const gTransform = computed(() => {
  const y = (props.wrapperRect.height - gridSize.height) / 2;
  const x = (props.wrapperRect.width - gridSize.width) / 2;

  return `translate(${x}, ${y})`;
});

function paint(event: MouseEvent) {
  const target = event.target as HTMLElement;
  const position = target.getAttribute('coord');
  if (!position) return;

  const color = event.buttons === 1 ? paletteStore.activeColor : null;
  beadsStore.paint(position as SchemaBeadCoord, color);
}

const listeners = computed(() => {
  return paletteStore.isPainting ? { mousemove: paint } : {};
});
</script>

<style scoped>
@layer page {
  .canvas-content:hover {
    cursor: crosshair;
  }
}
</style>
