<template>
  <g>
    <CanvasDefs>
      <rect
        :x
        :y
        :width
        :height
        :id="partialId"
        rx="2"
        ry="2"
        fill="none"
        stroke-width="1"
        class="canvas-selection"
      />
    </CanvasDefs>

    <use :href="partialHref" :stroke="emptyColor" />
    <use :href="partialHref" :stroke="selectionColor" stroke-dasharray="4 3" />
  </g>
</template>

<script setup lang="ts">
import { computed, useId } from 'vue';
import type { IBeadSelection } from '@editor/stores';
import {
  BEAD_SIZE,
  type IBeadsGrid,
  useBackgroundCanvasColor,
  useSelectionColor,
} from '@editor/composables';
import CanvasDefs from './CanvasDefs.vue';

const props = defineProps<{
  selected: IBeadSelection;
  beadsGrid: IBeadsGrid;
}>();

const fromOffset = computed(() => props.beadsGrid.resolveBeadOffset(props.selected.from));
const toOffset = computed(() => props.beadsGrid.resolveBeadOffset(props.selected.to));

const PADDING = 2;

const partialId = `editorSelection-${useId()}`;
const partialHref = `#${partialId}`;

const x = computed(() => Math.min(fromOffset.value[0], toOffset.value[0]) - PADDING);
const y = computed(() => Math.min(fromOffset.value[1], toOffset.value[1]) - PADDING);
const width = computed(() => Math.abs(fromOffset.value[0] - toOffset.value[0]) + BEAD_SIZE + PADDING * 2);
const height = computed(() => Math.abs(fromOffset.value[1] - toOffset.value[1]) + BEAD_SIZE + PADDING * 2);

const emptyColor = useBackgroundCanvasColor();
const selectionColor = useSelectionColor();
</script>

<style scoped>
@layer page {
  .canvas-selection {
    transition-property: x, y, width, height;
    transition-duration: 0.15s;
    transition-timing-function: ease-out;
    will-change: x, y, width, height;
  }
}
</style>
