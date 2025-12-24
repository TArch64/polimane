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
      <SelectionOverlay ref="overlayRef" v-if="isOverlayDisplaying" />

      <SelectionArea ref="areaRef" :selectionRef v-if="selectionRef">
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
import { AA_HIGH_CONTRAST, useBackgroundAccessibleColor } from '@editor/composables';
import { onClickOutside } from '@vueuse/core';
import { DirectionList } from '@/enums';
import { useDomRef } from '@/composables';
import ForeignTeleport from '../ForeignTeleport.vue';
import SelectionArea from './SelectionArea.vue';
import SelectionResizeHandle from './SelectionResizeHandle.vue';
import SelectionOverlay from './SelectionOverlay.vue';

const canvasStore = useCanvasStore();
const selectionStore = useSelectionStore();
const selectionColor = useBackgroundAccessibleColor(AA_HIGH_CONTRAST);

const selectionRef = ref<SVGElement>(null!);
const areaRef = useDomRef<HTMLElement | null>();
const overlayRef = useDomRef<HTMLElement | null>();

const isOverlayDisplaying = ref(false);

const clickOutside = onClickOutside(areaRef, selectionStore.reset, {
  ignore: [overlayRef],
  controls: true,
});

clickOutside.cancel();
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
