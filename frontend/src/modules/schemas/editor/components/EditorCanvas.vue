<template>
  <div
    ref="wrapperRef"
    class="editor-canvas"
    @contextmenu.prevent
    @keydown.esc="focusObjectStore.deactivatePath"
  >
    <KonvaStage
      :config
      :ref="onStageMounted"
      @wheel="onWheel"
      v-if="isReady"
    >
      <CanvasContent />
    </KonvaStage>
  </div>
</template>

<script setup lang="ts">
import { computed, nextTick, ref, type VNodeRef } from 'vue';
import { useElementSize } from '@vueuse/core';
import Konva from 'konva';
import type { KonvaStage } from 'vue-konva';
import type { KonvaEventObject } from 'konva/lib/Node';
import { useCanvasNavigation, useCanvasZoom, useMaybeCanvasStage } from '../composables';
import { useFocusObjectStore } from '../stores';
import { CanvasContent } from './content';

const canvasStage = useMaybeCanvasStage();
const focusObjectStore = useFocusObjectStore();

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
  const canvas: HTMLCanvasElement = stage?.content.querySelector('canvas');

  if (canvas) {
    canvas.tabIndex = 0;
    canvasStage.value = stage;
  }
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
