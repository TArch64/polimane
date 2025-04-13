<template>
  <div ref="wrapperRef" @contextmenu.prevent>
    <canvas ref="canvasRef" />
  </div>
</template>

<script setup lang="ts">
import { markRaw, onMounted, onUnmounted, type Ref, ref } from 'vue';
import { Canvas, Rect } from 'fabric';
import { provideCanvas, useCanvasNavigation, useCanvasZoom } from '../composables';

const canvasRef = ref<HTMLCanvasElement>(null!);
const wrapperRef = ref<HTMLElement>(null!);

const canvas: Ref<Canvas> = ref(null!);
provideCanvas(canvas);

onMounted(() => {
  canvas.value = markRaw(new Canvas(canvasRef.value, {
    selection: false,
    width: wrapperRef.value.offsetWidth,
    height: wrapperRef.value.offsetHeight,
  }));

  const rect = new Rect({
    selectable: false,
    hasControls: false,
    hasBorders: false,
    width: 10,
    height: 10,
    hoverCursor: 'default',
  });

  canvas.value.add(rect);
  canvas.value.centerObject(rect);
});

onUnmounted(() => {
  for (const object of canvas.value.getObjects()) {
    object.off();
  }

  canvas.value.destroy();
});

useCanvasZoom();
useCanvasNavigation();
</script>
