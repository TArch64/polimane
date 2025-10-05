<template>
  <main
    ref="wrapperRef"
    @contextmenu.prevent
    @keydown="onKeydown"
  >
    <svg
      ref="canvasRef"
      tabindex="0"
      class="editor-canvas__svg"
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
import { computed, nextTick, onMounted, reactive, ref } from 'vue';
import { useEditorStore } from '../stores';
import { useCanvasNavigation, useCanvasZoom } from '../composables';
import type { IViewBox } from '../types';
import { CanvasContent } from './content';

const editorStore = useEditorStore();

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

  await nextTick();
  canvasRef.value?.focus();
});

const canvasZoom = useCanvasZoom({ wrapperRect, viewBox });
const canvasNavigation = useCanvasNavigation({ canvasZoom, viewBox });

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
</script>

<style scoped>
@layer page {
  .editor-canvas__svg {
    outline: none !important;
  }
}
</style>
