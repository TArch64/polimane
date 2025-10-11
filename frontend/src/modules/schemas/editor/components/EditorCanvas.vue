<template>
  <main ref="wrapperRef" @contextmenu.prevent>
    <svg
      ref="canvasRef"
      xmlns="http://www.w3.org/2000/svg"
      preserveAspectRatio="xMidYMin slice"
      :width="wrapperRect.width"
      :height="wrapperRect.height"
      :viewBox="viewBoxAttr"
      @wheel="onWheel"
      v-if="wrapperRect"
    >
      <template v-if="canvasRef">
        <CanvasContent :wrapperRect />
        <CanvasSelection :canvasRef v-if="toolsStore.isSelection" />
      </template>
    </svg>
  </main>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue';
import { useHistoryStore, useToolsStore } from '../stores';
import { useCanvasNavigation, useCanvasZoom, useHotKeys } from '../composables';
import type { IViewBox } from '../types';
import { CanvasContent, CanvasSelection } from './content';

const historyStore = useHistoryStore();
const toolsStore = useToolsStore();

const canvasRef = ref<SVGSVGElement | null>(null);
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
const canvasNavigation = useCanvasNavigation({ viewBox });

function onWheel(event: WheelEvent): void {
  event.preventDefault();
  event.ctrlKey ? canvasZoom.zoom(event) : canvasNavigation.navigate(event);
}

useHotKeys({
  Meta_Z: () => historyStore.undo(),
  Meta_Shift_Z: () => historyStore.redo(),
});
</script>

<style scoped>
@layer page {
  :has(.canvas-selection):deep(.canvas-content) {
    cursor: default !important;
  }
}
</style>
