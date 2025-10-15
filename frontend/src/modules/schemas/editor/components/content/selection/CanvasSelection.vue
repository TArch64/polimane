<template>
  <g>
    <rect
      :x="selectionStore.area.x"
      :y="selectionStore.area.y"
      :width="selectionStore.area.width"
      :height="selectionStore.area.height"
      :stroke="selectionColor"
      rx="2"
      ry="2"
      fill="none"
      class="canvas-selection canvas-selection--original"
      stroke-dasharray="4 3"
      ref="selectionRef"
    />

    <rect
      :x="selectionStore.resize.x"
      :y="selectionStore.resize.y"
      :width="selectionStore.resize.width"
      :height="selectionStore.resize.height"
      :stroke="selectionColor"
      rx="2"
      ry="2"
      fill="none"
      class="canvas-selection canvas-selection--moving"
      stroke-opacity="0.5"
      v-if="selectionStore.resize.isResizing"
    />

    <ForeignTeleport>
      <SelectionOverlay v-if="isOverlayDisplaying" />

      <SelectionArea :selectionRef v-if="selectionRef">
        <SelectionResizeHandle
          v-for="direction of DirectionList"
          :direction
          :key="direction"
          v-model:overlay="isOverlayDisplaying"
        />
      </SelectionArea>
    </ForeignTeleport>
  </g>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useCanvasStore, useSelectionStore } from '@editor/stores';
import { useSelectionColor } from '@editor/composables';
import { DirectionList } from '@/enums';
import ForeignTeleport from '../ForeignTeleport.vue';
import SelectionArea from './SelectionArea.vue';
import SelectionResizeHandle from './SelectionResizeHandle.vue';
import SelectionOverlay from './SelectionOverlay.vue';

const canvasStore = useCanvasStore();
const selectionStore = useSelectionStore();
const selectionColor = useSelectionColor();

const selectionRef = ref<SVGElement>(null!);
const isOverlayDisplaying = ref(false);
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
}
</style>
