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
      @mousedown="paletteStore.setPainting(true)"
      @mouseup="paletteStore.setPainting(false)"
      v-if="isReady"
    >
      <CanvasContent />
    </KonvaStage>
  </main>
</template>

<script setup lang="ts">
import { computed, nextTick, ref, type VNodeRef } from 'vue';
import { useElementSize } from '@vueuse/core';
import Konva from 'konva';
import type { KonvaStage } from 'vue-konva';
import type { KonvaEventObject } from 'konva/lib/Node';
import { useCanvasNavigation, useCanvasZoom } from '../composables';
import { useFocusObjectStore, usePaletteStore } from '../stores';
import { CanvasContent } from './content';

const focusObjectStore = useFocusObjectStore();
const paletteStore = usePaletteStore();

const wrapperRef = ref<HTMLElement | null>(null);
const wrapperSize = useElementSize(wrapperRef);

const isReady = computed(() => !!wrapperSize.width.value && !!wrapperSize.height.value);

const config = computed((): Konva.StageConfig => ({
  width: wrapperSize.width.value,
  height: wrapperSize.height.value,
}));

const onStageMounted: VNodeRef = async (ref): Promise<void> => {
  await nextTick();
  const stage = (ref as InstanceType<KonvaStage>)?.getStage();
  window.__KONVA_STAGE_REF__.value = stage ?? null;
  const canvas = stage?.content.querySelector<HTMLCanvasElement>('canvas');
  if (canvas) canvas.tabIndex = 0;
};

const canvasZoom = useCanvasZoom();
const canvasNavigation = useCanvasNavigation();

function onWheel(event: KonvaEventObject<WheelEvent, Konva.Stage>): void {
  event.evt.preventDefault();
  event.evt.ctrlKey ? canvasZoom.zoom(event) : canvasNavigation.navigate(event);
  event.currentTarget.batchDraw();
}
</script>

<style scoped>
.editor-canvas:deep(canvas) {
  outline: none;
}
</style>
