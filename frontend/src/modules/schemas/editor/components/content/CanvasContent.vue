<template>
  <CanvasDefs>
    <CanvasBackgroundPattern :id="backgroundPatternId" />
  </CanvasDefs>

  <g :transform class="canvas-content" v-on="listeners">
    <rect
      ref="backgroundRef"
      :fill="backgroundPatternFill"
      :x="beadsGrid.size.minX"
      :y="beadsGrid.size.minY"
      :width="beadsGrid.size.width"
      :height="beadsGrid.size.height"
    />

    <CanvasBead
      v-for="bead of beadsGrid.beads"
      :bead
      :key="bead.coord"
    />

    <FadeTransition>
      <CanvasSelection
        :key="`${selected.from}-${selected.to}`"
        v-if="selected"
      />
    </FadeTransition>
  </g>
</template>

<script setup lang="ts">
import { computed, ref, useId } from 'vue';
import { useSelectionStore } from '@editor/stores';
import { FadeTransition } from '@/components/transition';
import { type IBeadsGrid, useBeadTools } from '../../composables';
import { CanvasBead } from './CanvasBead';
import CanvasBackgroundPattern from './CanvasBackgroundPattern.vue';
import { CanvasSelection } from './selection';
import CanvasDefs from './CanvasDefs.vue';

const props = defineProps<{
  wrapperRect: DOMRect;
  beadsGrid: IBeadsGrid;
}>();

const selectionStore = useSelectionStore();
const selected = computed(() => selectionStore.selected);

const backgroundPatternId = `editorEmptyBeads-${useId()}`;
const backgroundPatternFill = `url(#${backgroundPatternId})`;
const backgroundRef = ref<SVGRectElement>(null!);

const listeners = useBeadTools({
  backgroundRef,
  beadsGrid: props.beadsGrid,
});

const transform = (() => {
  // shouldn't be recomputed to avoid shifting on schema resize
  const y = (props.wrapperRect.height - props.beadsGrid.size.height) / 2;
  const x = (props.wrapperRect.width - props.beadsGrid.size.width) / 2;

  return `translate(${x}, ${y})`;
})();
</script>

<style scoped>
@layer page {
  .canvas-content:hover {
    cursor: var(--editor-cursor, crosshair);
  }
}
</style>
