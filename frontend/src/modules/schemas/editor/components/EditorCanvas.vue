<template>
  <main
    ref="wrapperRef"
    @contextmenu.prevent
    @keydown="onKeydown"
  >
    <svg
      :viewBox
      ref="canvasRef"
      tabindex="0"
      class="editor-canvas__svg"
      xmlns="http://www.w3.org/2000/svg"
      preserveAspectRatio="xMidYMin slice"
      :width="wrapperRect.width"
      :height="wrapperRect.height"
      @wheel="onWheel"
      @mousedown="togglePainting"
      @mouseup="togglePainting"
      v-if="wrapperRect"
    >
      <CanvasContent :wrapperRect />
    </svg>
  </main>
</template>

<script setup lang="ts">
import { computed, nextTick, onMounted, ref } from 'vue';
import { useEditorStore, usePaletteStore } from '../stores';
import { useCanvasNavigation, useCanvasZoom } from '../composables';
import { CanvasContent } from './content';

const editorStore = useEditorStore();
const paletteStore = usePaletteStore();

const canvasRef = ref<SVGSVGElement | null>(null);
const wrapperRef = ref<HTMLElement | null>(null);
const wrapperRect = ref<DOMRect | null>(null);

const viewBox = computed(() => [
  0,
  0,
  wrapperRect.value!.width,
  wrapperRect.value!.height,
].join(' '));

onMounted(async () => {
  wrapperRect.value = wrapperRef.value!.getBoundingClientRect();
  await nextTick();
  canvasRef.value!.focus();
});

const canvasZoom = useCanvasZoom({ wrapperRect });
const canvasNavigation = useCanvasNavigation(canvasZoom);

function onWheel(event: WheelEvent): void {
  event.preventDefault();
  event.ctrlKey ? canvasZoom.zoom(event) : canvasNavigation.navigate(event);
}

function onKeydown(event: KeyboardEvent) {
  if (!event.metaKey || event.key.toLowerCase() !== 'z') {
    return;
  }

  event.preventDefault();
  event.shiftKey ? editorStore.redo() : editorStore.undo();
}

function togglePainting(event: MouseEvent) {
  if (event.buttons > 1) return;
  paletteStore.setPainting(event.buttons === 1);
}
</script>

<style scoped>
@layer page {
  .editor-canvas__svg {
    outline: none !important;
  }
}
</style>
