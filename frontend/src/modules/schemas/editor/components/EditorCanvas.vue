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
      v-if="isReady"
    >
      <KonvaLayer ref="layerRef" :config="layerConfig">
        <CanvasContent :stage-config="config" />
      </KonvaLayer>
    </KonvaStage>
  </main>
</template>

<script setup lang="ts">
import { computed, nextTick, ref, watch } from 'vue';
import { useDebounceFn, useElementSize } from '@vueuse/core';
import Konva from 'konva';
import type { KonvaEventObject } from 'konva/lib/Node';
import { provideCanvasStage, useCanvasNavigation, useCanvasZoom, useNodeRef } from '../composables';
import { useEditorStore, usePaletteStore } from '../stores';
import { CanvasContent } from './content';

const editorStore = useEditorStore();
const paletteStore = usePaletteStore();

const wrapperRef = ref<HTMLElement | null>(null);
const { width: canvasWidth, height: canvasHeight } = useElementSize(wrapperRef);
const isReady = computed(() => !!canvasWidth.value && !!canvasHeight.value);

const stageRef = useNodeRef<Konva.Stage>();
const layerRef = useNodeRef<Konva.Layer>();

provideCanvasStage(stageRef);

const config = computed(() => ({
  width: canvasWidth.value,
  height: canvasHeight.value,
}));

watch(stageRef, async (stage) => {
  if (stage) {
    stage.content.querySelector<HTMLElement>('canvas')!.tabIndex = 0;

    await nextTick();

    layerRef.value.to({
      opacity: 1,
      duration: 0.3,
      easing: Konva.Easings.EaseOut,
    });
  }
});

const layerConfig: Konva.LayerConfig = {
  id: 'editor-layer',
  opacity: 0,
};

const canvasZoom = useCanvasZoom();
const canvasNavigation = useCanvasNavigation();

let isLayerCachingEnabled = false;

function enableLayerCaching(): void {
  if (!isLayerCachingEnabled) {
    layerRef.value.cache();
    isLayerCachingEnabled = true;
  }
}

const disableLayerCaching = useDebounceFn(() => {
  layerRef.value.clearCache();
  isLayerCachingEnabled = false;
}, 100);

function onWheel(event: KonvaEventObject<WheelEvent, Konva.Stage>): void {
  enableLayerCaching();
  event.evt.preventDefault();
  event.evt.ctrlKey ? canvasZoom.zoom(event) : canvasNavigation.navigate(event);
  event.currentTarget.batchDraw();
  disableLayerCaching();
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
