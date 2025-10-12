<template>
  <g>
    <rect
      :x
      :y
      :width
      :height
      :stroke="selectionColor"
      rx="2"
      ry="2"
      fill="none"
      class="canvas-selection canvas-selection--original"
      stroke-dasharray="4 3"
      ref="selectionRef"
    />

    <rect
      :x="movingX"
      :y="movingY"
      :width="movingWidth"
      :height="movingHeight"
      :stroke="selectionColor"
      rx="2"
      ry="2"
      fill="none"
      class="canvas-selection canvas-selection--moving"
      stroke-opacity="0.5"
      v-if="isMoving"
    />

    <ForeignTeleport>
      <div
        class="canvas-selection__overlay"
        v-if="isOverlayDisplaying"
      />

      <SelectionArea :selectionRef v-if="selectionRef">
        <SelectionResizeHandle
          v-for="direction of DirectionList"
          :direction
          :key="direction"
          v-model:translation="translation[direction]"
          v-model:overlay="isOverlayDisplaying"
        />
      </SelectionArea>
    </ForeignTeleport>
  </g>
</template>

<script setup lang="ts">
import { computed, reactive, ref } from 'vue';
import { type IBeadSelection, useCanvasStore } from '@editor/stores';
import { BEAD_SIZE, type IBeadsGrid, useSelectionColor } from '@editor/composables';
import { Direction, DirectionList } from '@/enums';
import ForeignTeleport from '../ForeignTeleport.vue';
import SelectionArea from './SelectionArea.vue';
import SelectionResizeHandle from './SelectionResizeHandle.vue';

const props = defineProps<{
  selected: IBeadSelection;
  beadsGrid: IBeadsGrid;
}>();

const canvasStore = useCanvasStore();
const selectionColor = useSelectionColor();

const selectionRef = ref<SVGElement>(null!);
const isOverlayDisplaying = ref(false);

const translation = reactive<Record<Direction, number>>({
  [Direction.TOP]: 0,
  [Direction.BOTTOM]: 0,
  [Direction.LEFT]: 0,
  [Direction.RIGHT]: 0,
});

const isMoving = computed(() => {
  return Object.values(translation).some((value) => value !== 0);
});

const fromOffset = props.beadsGrid.resolveBeadOffset(props.selected.from);
const toOffset = props.beadsGrid.resolveBeadOffset(props.selected.to);

const PADDING = 2;
const x = Math.min(fromOffset.x, toOffset.x) - PADDING;
const y = Math.min(fromOffset.y, toOffset.y) - PADDING;
const width = Math.abs(fromOffset.x - toOffset.x) + BEAD_SIZE + PADDING * 2;
const height = Math.abs(fromOffset.y - toOffset.y) + BEAD_SIZE + PADDING * 2;

const movingX = computed(() => x - translation.left);
const movingY = computed(() => y - translation.top);
const movingWidth = computed(() => width + translation.right + translation.left);
const movingHeight = computed(() => height + translation.bottom + translation.top);
</script>

<style scoped>
@layer page {
  .canvas-selection {
    stroke-width: calc(var(--selection-size) + (1 / v-bind("canvasStore.scale")));
    will-change: stroke-width;
  }

  .canvas-selection--original {
    --selection-size: 1
  }

  .canvas-selection--moving {
    --selection-size: 0
  }

  .canvas-selection__overlay {
    position: fixed;
    z-index: 999;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
  }
}
</style>
