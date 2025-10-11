<template>
  <main ref="wrapperRef" :class="wrapperClasses" @contextmenu.prevent>
    <svg
      ref="canvasRef"
      xmlns="http://www.w3.org/2000/svg"
      preserveAspectRatio="xMidYMin slice"
      :width="wrapperRect.width"
      :height="wrapperRect.height"
      :viewBox
      @wheel="onWheel"
      v-if="wrapperRect"
    >
      <defs id="editorCanvasDefs" />

      <CanvasContent
        :beadsGrid
        :wrapperRect
        v-if="canvasRef"
      />
    </svg>

    <Teleport to="body">
      <FadeTransition @after-leave="selectionStore.reset">
        <EditorSelection v-if="toolsStore.isSelection && selectionStore.isSelecting" />
      </FadeTransition>
    </Teleport>
  </main>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { FadeTransition } from '@/components/transition';
import { useCanvasStore, useHistoryStore, useSelectionStore, useToolsStore } from '../stores';
import { useBeadsGrid, useCanvasNavigation, useCanvasZoom, useHotKeys } from '../composables';
import { CanvasContent } from './content';
import EditorSelection from './EditorSelection.vue';

const historyStore = useHistoryStore();
const toolsStore = useToolsStore();
const selectionStore = useSelectionStore();
const canvasStore = useCanvasStore();

const canvasRef = ref<SVGSVGElement | null>(null);
const wrapperRef = ref<HTMLElement | null>(null);
const wrapperRect = ref<DOMRect | null>(null);

const wrapperClasses = computed(() => ({
  'canvas-editor--selection': toolsStore.isSelection,
}));

onMounted(() => {
  wrapperRect.value = wrapperRef.value!.getBoundingClientRect();
});

const viewBox = computed((): string => {
  const { x, y } = canvasStore.translation;
  const width = wrapperRect.value!.width / canvasStore.scale;
  const height = wrapperRect.value!.height / canvasStore.scale;
  return `${x} ${y} ${width} ${height}`;
});

const canvasZoom = useCanvasZoom();
const canvasNavigation = useCanvasNavigation();
const beadsGrid = useBeadsGrid();

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
  .canvas-editor--selection {
    --editor-cursor: default;
  }
}
</style>
