<template>
  <div ref="wrapperRef">
    <canvas ref="canvasRef" />
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { Canvas, Rect } from 'fabric';

const canvasRef = ref<HTMLCanvasElement>(null!);
const wrapperRef = ref<HTMLElement>(null!);

let canvas: Canvas;

onMounted(() => {
  canvas = new Canvas(canvasRef.value, {
    selection: false,
    width: wrapperRef.value.offsetWidth,
    height: wrapperRef.value.offsetHeight,
    centeredScaling: true,
  });

  const rect = new Rect({
    selectable: false,
    hasControls: false,
    hasBorders: false,
    width: 10,
    height: 10,
    hoverCursor: 'default',
  });

  canvas.add(rect);
  canvas.centerObject(rect);
});
</script>
