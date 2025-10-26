<template>
  <main ref="wrapperRef" :class="wrapperClasses" @contextmenu.prevent>
    <svg
      ref="canvasRef"
      xmlns="http://www.w3.org/2000/svg"
      preserveAspectRatio="xMidYMin slice"
      class="canvas-editor"
      :width="wrapperRect.width"
      :height="wrapperRect.height"
      :viewBox
      v-on="canvasEvents.listeners"
      v-if="wrapperRect"
    >
      <defs id="editorCanvasDefs" />

      <CanvasContent
        :beadsGrid
        :wrapperRect
        v-if="canvasRef"
      />
    </svg>

    <Teleport to="body" v-if="editorStore.canEdit && !isMobile">
      <FadeTransition>
        <EditorSelection
          v-if="toolsStore.isSelection && selectionStore.isSelecting"
        />
      </FadeTransition>
    </Teleport>
  </main>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { FadeTransition } from '@/components/transition';
import { useMobileScreen } from '@/composables';
import { useCanvasStore, useEditorStore, useSelectionStore, useToolsStore } from '../stores';
import { useBeadsGrid, useCanvasEvents, useHotKeys } from '../composables';
import { CanvasContent } from './content';
import EditorSelection from './EditorSelection.vue';

const editorStore = useEditorStore();
const toolsStore = useToolsStore();
const selectionStore = useSelectionStore();
const canvasStore = useCanvasStore();

const isMobile = useMobileScreen();

const canvasRef = ref<SVGSVGElement | null>(null);
const wrapperRef = ref<HTMLElement | null>(null);
const wrapperRect = ref<DOMRect | null>(null);

const wrapperClasses = computed(() => ({
  'canvas-editor--selection': toolsStore.isSelection,
  'canvas-editor--readonly': !editorStore.canEdit,
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

const canvasEvents = useCanvasEvents();
const beadsGrid = useBeadsGrid();

useHotKeys({
  Backspace: selectionStore.removeSelected,
  Delete: selectionStore.removeSelected,
}, {
  isActive: () => toolsStore.isSelection,
});
</script>

<style scoped>
@layer page {
  .canvas-editor--readonly,
  .canvas-editor--selection {
    --editor-cursor: default;
  }
}
</style>
