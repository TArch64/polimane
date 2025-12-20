<template>
  <g :transform ref="groupRef" class="canvas-content">
    <template v-if="editorStore.canEdit">
      <rect
        ref="backgroundRef"
        fill="transparent"
        v-bind="backgroundRect"
      />

      <BackgroundLinear :rect="backgroundRect" />
    </template>

    <CanvasBead
      v-for="item of beadsGrid.beads"
      :key="item.coord"
      :item
    />

    <FadeTransition v-if="editorStore.canEdit && !isMobile">
      <CanvasSelection
        :key="`${selected.from}-${selected.to}`"
        v-if="selected"
      />
    </FadeTransition>
  </g>
</template>

<script setup lang="ts">
import { computed, ref, toRef } from 'vue';
import { useEditorStore, useSelectionStore } from '@editor/stores';
import { BEAD_BUGLE_CORNER_RADIUS } from '@editor/const';
import { FadeTransition } from '@/components/transition';
import { useMobileScreen } from '@/composables';
import type { INodeRect } from '@/models';
import { type IBeadsGrid, useEditorTools } from '../../composables';
import { CanvasBead } from './CanvasBead';
import { BackgroundLinear } from './background';
import { CanvasSelection } from './selection';

const props = defineProps<{
  canvasRef: SVGSVGElement;
  wrapperRect: DOMRect;
  beadsGrid: IBeadsGrid;
}>();

const CORNER_RADIUS = BEAD_BUGLE_CORNER_RADIUS;

const editorStore = useEditorStore();
const selectionStore = useSelectionStore();
const selected = computed(() => selectionStore.selected);

const isMobile = useMobileScreen();

const backgroundRef = ref<SVGRectElement>(null!);
const groupRef = ref<SVGGElement>(null!);

const backgroundRect = computed((): INodeRect => ({
  x: props.beadsGrid.size.minX,
  y: props.beadsGrid.size.minY,
  width: props.beadsGrid.size.width,
  height: props.beadsGrid.size.height,
}));

if (editorStore.canEdit && !isMobile.value) {
  useEditorTools({
    groupRef,
    canvasRef: toRef(props, 'canvasRef'),
    backgroundRef,
    beadsGrid: props.beadsGrid,
  });
}

const transform = (() => {
  // shouldn't be recomputed to avoid shifting on schema resize
  const y = (props.wrapperRect.height - props.beadsGrid.size.height) / 2;
  const x = (props.wrapperRect.width - props.beadsGrid.size.width) / 2;

  return `translate(${x}, ${y})`;
})();
</script>

<style scoped>
@layer page {
  :deep(.canvas-bead-bugle),
  :deep(.canvas-bead-ref) {
    rx: v-bind("CORNER_RADIUS");
    ry: v-bind("CORNER_RADIUS");
  }

  :deep(.canvas-bead-bugle) {
    transition: 0.15s ease-out;
    transition-property: width, height, x, y;
    will-change: width, height, x, y;
  }
}
</style>
