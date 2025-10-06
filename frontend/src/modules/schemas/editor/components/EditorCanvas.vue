<template>
  <main ref="wrapperRef" @contextmenu.prevent>
    <svg
      xmlns="http://www.w3.org/2000/svg"
      preserveAspectRatio="xMidYMin slice"
      :width="wrapperRect.width"
      :height="wrapperRect.height"
      :viewBox="viewBoxAttr"
      @wheel="onWheel"
      v-if="wrapperRect"
    >
      <CanvasContent
        :canvasZoom
        :wrapperRect
      />
    </svg>
  </main>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue';
import { useHistoryStore } from '../stores';
import { useCanvasNavigation, useCanvasZoom, useHotKeys } from '../composables';
import type { IViewBox } from '../types';
import { CanvasContent } from './content';

const historyStore = useHistoryStore();

const wrapperRef = ref<HTMLElement | null>(null);
const wrapperRect = ref<DOMRect | null>(null);

const viewBox = reactive<IViewBox>({
  x: 0,
  y: 0,
  width: 0,
  height: 0,
});

const viewBoxAttr = computed(() => {
  const { height, width, x, y } = viewBox;
  return `${x} ${y} ${width} ${height}`;
});

onMounted(async () => {
  wrapperRect.value = wrapperRef.value!.getBoundingClientRect();
  viewBox.width = wrapperRect.value.width;
  viewBox.height = wrapperRect.value.height;
});

const canvasZoom = useCanvasZoom({ wrapperRect, viewBox });
const canvasNavigation = useCanvasNavigation({ canvasZoom, viewBox });

function onWheel(event: WheelEvent): void {
  event.preventDefault();
  event.ctrlKey ? canvasZoom.zoom(event) : canvasNavigation.navigate(event);
}

useHotKeys({
  Meta_Z: () => historyStore.undo(),
  Meta_Shift_Z: () => historyStore.redo(),
});
</script>
