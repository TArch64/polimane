<template>
  <main ref="wrapperRef" @contextmenu.prevent>
    <svg
      ref="canvasRef"
      xmlns="http://www.w3.org/2000/svg"
      preserveAspectRatio="xMidYMin slice"
      tabindex="0"
      class="canvas-editor"
      :class="canvasClasses"
      :width="wrapperRect.width"
      :height="wrapperRect.height"
      :viewBox
      v-on="canvasEvents.listeners"
      v-if="wrapperRect"
    >
      <defs id="editorCanvasDefs" />

      <CanvasContent
        :canvasRef
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
import { EditorCursor } from '@editor/enums';
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

const cursor = computed(() => {
  return editorStore.canEdit ? canvasStore.cursor : EditorCursor.DEFAULT;
});

const canvasClasses = computed(() => [
  `canvas-editor--cursor-${canvasStore.cursorTarget}`,
]);

onMounted(() => {
  wrapperRect.value = wrapperRef.value!.getBoundingClientRect();
  canvasRef.value?.focus();
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
  .canvas-editor {
    outline: none;
    user-select: none;
    --editor-cursor: v-bind("cursor");
  }

  .canvas-editor--cursor-canvas {
    cursor: var(--editor-cursor);
  }

  .canvas-editor--cursor-content:deep(.canvas-content) {
    cursor: var(--editor-cursor);
  }
}
</style>
