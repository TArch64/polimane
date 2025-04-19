<template>
  <div ref="wrapperRef" @contextmenu.prevent>
    <canvas ref="canvasRef" />
  </div>
</template>

<script setup lang="ts">
import { markRaw, onMounted, onUnmounted, type Ref, ref } from 'vue';
import { Canvas } from 'fabric';
import {
  provideCanvas,
  useCanvasContent,
  useCanvasNavigation,
  useCanvasZoom,
} from '../composables';

const canvasRef = ref<HTMLCanvasElement>(null!);
const wrapperRef = ref<HTMLElement>(null!);
const canvas: Ref<Canvas | null> = ref(null);

provideCanvas(canvas);

onMounted(() => {
  canvas.value = markRaw(new Canvas(canvasRef.value, {
    selection: false,
    // TODO update navigation to work correctly with enabled skipOffscreen
    skipOffscreen: false,
    width: wrapperRef.value.offsetWidth,
    height: wrapperRef.value.offsetHeight,
  }));
});

onUnmounted(() => {
  for (const object of canvas.value!.getObjects()) {
    object.off();
  }

  canvas.value!.destroy();
});

useCanvasZoom();
useCanvasNavigation();
useCanvasContent();
</script>
