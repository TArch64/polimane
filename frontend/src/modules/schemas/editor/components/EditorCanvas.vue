<template>
  <main
    ref="wrapperRef"
    class="editor-canvas"
    @contextmenu.prevent
    @keydown.esc="focusObjectStore.deactivatePath"
  >
    <KonvaStage
      :config
      :ref="onStageMounted"
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
import { computed, nextTick, ref, type VNodeRef } from 'vue';
import { useDebounceFn, useElementSize } from '@vueuse/core';
import Konva from 'konva';
import type { KonvaStage } from 'vue-konva';
import type { KonvaEventObject } from 'konva/lib/Node';
import { useCanvasNavigation, useCanvasZoom } from '../composables';
import { useFocusObjectStore, usePaletteStore } from '../stores';
import { CanvasContent, type IGroupLayoutEvent } from './content';

const focusObjectStore = useFocusObjectStore();
const paletteStore = usePaletteStore();

const wrapperRef = ref<HTMLElement | null>(null);
const wrapperSize = useElementSize(wrapperRef);

const isReady = computed(() => !!wrapperSize.width.value && !!wrapperSize.height.value);

const setRendered = useDebounceFn(async (event: Konva.KonvaEventObject<IGroupLayoutEvent>) => {
  const stage = event.currentTarget as Konva.Stage;

  stage.off('layout', setRendered);

  stage.findOne(`#${layerConfig.id}`)!.to({
    opacity: 1,
    duration: 0.1,
    easing: Konva.Easings.EaseOut,
  });
}, 100);

const config = computed(() => ({
  width: wrapperSize.width.value,
  height: wrapperSize.height.value,
}));

const onStageMounted: VNodeRef = async (ref): Promise<void> => {
  await nextTick();
  const stage = (ref as InstanceType<KonvaStage>)?.getStage();

  if (!stage) {
    window.__KONVA_STAGE_REF__.value = null;
    return;
  }

  window.__KONVA_STAGE_REF__.value = stage;
  stage.content.querySelector<HTMLElement>('canvas')!.tabIndex = 0;
  stage.on('layout', setRendered);
};

const layerConfig: Konva.LayerConfig = {
  id: 'editor-layer',
  opacity: 0,
};

const canvasZoom = useCanvasZoom();
const canvasNavigation = useCanvasNavigation();

function onWheel(event: KonvaEventObject<WheelEvent, Konva.Stage>): void {
  event.evt.preventDefault();
  event.evt.ctrlKey ? canvasZoom.zoom(event) : canvasNavigation.navigate(event);
  event.currentTarget.batchDraw();
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
