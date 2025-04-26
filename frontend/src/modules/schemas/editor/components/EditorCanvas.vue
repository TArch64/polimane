<template>
  <div ref="wrapperRef" @contextmenu.prevent>
    <canvas ref="canvasRef" />
    <CanvasSchema v-if="canvas" />
  </div>
</template>

<script setup lang="ts">
import { Canvas } from 'fabric';
import { markRaw, onMounted, onUnmounted, type Ref, ref } from 'vue';
import { provideCanvas, useCanvasNavigation, useCanvasZoom } from '../composables';
import { CanvasSchema } from './content';

const canvasRef = ref<HTMLCanvasElement>(null!);
const wrapperRef = ref<HTMLElement>(null!);
const canvas: Ref<Canvas | null> = ref(null);

provideCanvas(canvas);

onMounted(() => {
  canvas.value = markRaw(new Canvas(canvasRef.value, {
    selection: false,
    fireRightClick: true,
    width: wrapperRef.value.offsetWidth,
    height: wrapperRef.value.offsetHeight,
    hoverCursor: 'default',
  }));
});

onUnmounted(() => canvas.value!.dispose());

useCanvasZoom();
useCanvasNavigation();

if (import.meta.hot) {
  import.meta.hot.on('vite:beforeUpdate', () => {
    window.location.reload();
  });
}
</script>
