<template>
  <main
    ref="wrapperRef"
    class="editor-canvas"
    @contextmenu.prevent
    @keydown="onKeydown"
  >
    <KonvaStage
      :config
      ref="stageRef"
      @wheel="onWheel"
      @mousedown="togglePainting"
      @mouseup="togglePainting"
      @layout="setRendered"
      v-if="isReady"
    >
      <KonvaLayer :config="layerConfig">
        <CanvasContent />
      </KonvaLayer>
    </KonvaStage>
  </main>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { useDebounceFn, useElementSize } from '@vueuse/core';
import Konva from 'konva';
import type { KonvaEventObject } from 'konva/lib/Node';
import {
  provideNodeContextMenu,
  useCanvasCursor,
  useCanvasNavigation,
  useCanvasZoom,
  useNodeRef,
} from '../composables';
import { useEditorStore, usePaletteStore } from '../stores';
import { CanvasContent, type IGroupLayoutEvent } from './content';

const editorStore = useEditorStore();
const paletteStore = usePaletteStore();

const wrapperRef = ref<HTMLElement | null>(null);
const wrapperSize = useElementSize(wrapperRef);

const isReady = computed(() => !!wrapperSize.width.value && !!wrapperSize.height.value);

const setRendered = useDebounceFn(async (event: Konva.KonvaEventObject<IGroupLayoutEvent>) => {
  const stage = event.currentTarget as Konva.Stage;

  stage.off('layout', setRendered);

  stage.findOne(`#${layerConfig.id}`)!.to({
    opacity: 1,
    duration: 0.3,
    easing: Konva.Easings.EaseOut,
  });
}, 100);

const stageRef = useNodeRef<Konva.Stage>();

const config = computed(() => ({
  width: wrapperSize.width.value,
  height: wrapperSize.height.value,
}));

watch(stageRef, async (stage) => {
  window.__KONVA_STAGE_REF__.value = stage;

  if (stage) {
    stage.content.querySelector<HTMLElement>('canvas')!.tabIndex = 0;
    stage.on('layout', setRendered);
  }
});

const layerConfig: Konva.LayerConfig = {
  id: 'editor-layer',
  opacity: 0,
};

provideNodeContextMenu(stageRef);
const canvasZoom = useCanvasZoom();
const canvasNavigation = useCanvasNavigation();
useCanvasCursor(stageRef);

function onWheel(event: KonvaEventObject<WheelEvent, Konva.Stage>): void {
  event.evt.preventDefault();
  event.evt.ctrlKey ? canvasZoom.zoom(event) : canvasNavigation.navigate(event);
  event.currentTarget.batchDraw();
}

function onKeydown(event: KeyboardEvent) {
  if (!event.metaKey || event.key.toLowerCase() !== 'z') {
    return;
  }

  event.preventDefault();
  event.shiftKey ? editorStore.redo() : editorStore.undo();
}

function togglePainting(event: Konva.KonvaEventObject<MouseEvent>) {
  if (event.evt.buttons > 1) return;
  paletteStore.setPainting(event.evt.buttons === 1);
}
</script>

<style scoped>
.editor-canvas:deep(canvas) {
  outline: none;
}
</style>
