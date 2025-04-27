<template>
  <div ref="wrapperRef" @contextmenu.prevent>
    <KonvaStage
      :config
      @wheel="onWheel"
      v-if="isReady"
    >
      <CanvasContent />
    </KonvaStage>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import { useElementSize } from '@vueuse/core';
import Konva from 'konva';
import type { KonvaEventObject } from 'konva/lib/Node';
import { CanvasContent } from './content';

const wrapperRef = ref<HTMLElement | null>(null);
const wrapperSize = useElementSize(wrapperRef);

const isReady = computed(() => !!wrapperSize.width.value && !!wrapperSize.height.value);

const config = computed((): Konva.StageConfig => ({
  width: wrapperSize.width.value,
  height: wrapperSize.height.value,
}));

const MIN_ZOOM = 0.5;
const MAX_ZOOM = 10;

function pinchZoom(stage: Konva.Stage, event: WheelEvent): void {
  const oldScale = stage.scaleX();
  const pointer = stage.getPointerPosition()!;

  const mousePointTo = {
    x: (pointer.x - stage.x()) / oldScale,
    y: (pointer.y - stage.y()) / oldScale,
  };

  const scaleFactor = 1 - event.deltaY * 0.01;
  const newScale = Math.min(Math.max(oldScale * scaleFactor, MIN_ZOOM), MAX_ZOOM);

  stage.scale({
    x: newScale,
    y: newScale,
  });

  const newPos = {
    x: pointer.x - mousePointTo.x * newScale,
    y: pointer.y - mousePointTo.y * newScale,
  };

  stage.position(newPos);
}

function navigate(stage: Konva.Stage, event: WheelEvent) {
  const dx = event.deltaX;
  const dy = event.deltaY;

  const currentPos = stage.position();

  stage.position({
    x: currentPos.x - dx,
    y: currentPos.y - dy,
  });
}

function onWheel(kEvent: KonvaEventObject<WheelEvent, Konva.Stage>): void {
  const { evt: event, currentTarget: stage } = kEvent;
  event.preventDefault();
  event.ctrlKey ? pinchZoom(stage, event) : navigate(stage, event);
  stage.batchDraw();
}
</script>
