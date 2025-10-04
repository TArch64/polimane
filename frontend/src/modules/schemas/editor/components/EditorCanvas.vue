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
import { nextTick, onMounted, ref } from 'vue';
import { useEditorStore } from '../stores';
import { useCanvasNavigation, useCanvasZoom } from '../composables';
import { CanvasContent } from './content';

const editorStore = useEditorStore();

const canvasRef = ref<SVGSVGElement | null>(null);

const wrapperRef = ref<HTMLElement | null>(null);
const wrapperRect = ref<DOMRect | null>(null);

const viewBox = ref('');

onMounted(async () => {
  wrapperRect.value = wrapperRef.value!.getBoundingClientRect();

  viewBox.value = [
    0,
    0,
    wrapperRect.value.width,
    wrapperRect.value.height,
  ].join(' ');

  await nextTick();
  canvasRef.value?.focus();
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
</script>

<style scoped>
@layer page {
  .editor-canvas__svg {
    outline: none !important;
  }
}
</style>
