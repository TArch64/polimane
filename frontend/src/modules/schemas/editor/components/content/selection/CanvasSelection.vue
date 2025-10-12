<template>
  <g>
    <rect
      :x
      :y
      :width
      :height
      rx="2"
      ry="2"
      fill="none"
      ref="selectionRef"
      class="canvas-selection"
      :stroke="selectionColor"
      stroke-dasharray="4 3"
    />

    <ForeignTeleport>
      <SelectionArea :selectionRef v-if="selectionRef">
        <SelectionResizeHandle position="top" />
        <SelectionResizeHandle position="bottom" />
        <SelectionResizeHandle position="left" />
        <SelectionResizeHandle position="right" />
      </SelectionArea>
    </ForeignTeleport>
  </g>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { type IBeadSelection, useCanvasStore } from '@editor/stores';
import { BEAD_SIZE, type IBeadsGrid, useSelectionColor } from '@editor/composables';
import ForeignTeleport from '../ForeignTeleport.vue';
import SelectionArea from './SelectionArea.vue';
import SelectionResizeHandle from './SelectionResizeHandle.vue';

const props = defineProps<{
  selected: IBeadSelection;
  beadsGrid: IBeadsGrid;
}>();

const canvasStore = useCanvasStore();

const selectionRef = ref<SVGElement>(null!);

const fromOffset = props.beadsGrid.resolveBeadOffset(props.selected.from);
const toOffset = props.beadsGrid.resolveBeadOffset(props.selected.to);

const PADDING = 2;
const x = Math.min(fromOffset.x, toOffset.x) - PADDING;
const y = Math.min(fromOffset.y, toOffset.y) - PADDING;
const width = Math.abs(fromOffset.x - toOffset.x) + BEAD_SIZE + PADDING * 2;
const height = Math.abs(fromOffset.y - toOffset.y) + BEAD_SIZE + PADDING * 2;

const selectionColor = useSelectionColor();
</script>

<style scoped>
@layer page {
  .canvas-selection {
    stroke-width: calc(1 + (1 / v-bind("canvasStore.scale")));
    will-change: stroke-width;
  }
}
</style>
