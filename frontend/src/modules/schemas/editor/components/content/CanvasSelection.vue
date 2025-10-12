<template>
  <rect
    :x
    :y
    :width
    :height
    rx="2"
    ry="2"
    fill="none"
    class="canvas-selection"
    :stroke="selectionColor"
    stroke-dasharray="4 3"
  />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { type IBeadSelection, useCanvasStore } from '@editor/stores';
import { BEAD_SIZE, type IBeadsGrid, useSelectionColor } from '@editor/composables';

const props = defineProps<{
  selected: IBeadSelection;
  beadsGrid: IBeadsGrid;
}>();

const canvasStore = useCanvasStore();

const fromOffset = computed(() => props.beadsGrid.resolveBeadOffset(props.selected.from));
const toOffset = computed(() => props.beadsGrid.resolveBeadOffset(props.selected.to));

const PADDING = 2;

const x = computed(() => Math.min(fromOffset.value.x, toOffset.value.x) - PADDING);
const y = computed(() => Math.min(fromOffset.value.y, toOffset.value.y) - PADDING);
const width = computed(() => Math.abs(fromOffset.value.x - toOffset.value.x) + BEAD_SIZE + PADDING * 2);
const height = computed(() => Math.abs(fromOffset.value.y - toOffset.value.y) + BEAD_SIZE + PADDING * 2);

const selectionColor = useSelectionColor();
</script>

<style scoped>
@layer page {
  .canvas-selection {
      stroke-width: calc(2 / v-bind("canvasStore.scale"));
    transition-property: x, y, width, height;
    transition-duration: 0.15s;
    transition-timing-function: ease-out;
    will-change: x, y, width, height;
  }
}
</style>
