<template>
  <KonvaRect ref="rectRef" :config />
</template>

<script setup lang="ts">
import Konva from 'konva';
import { computed } from 'vue';
import { useNodeClientRect, useNodeParent, useNodeRef } from '@/modules/schemas/editor/composables';

const rectRef = useNodeRef<Konva.Rect | null>();
const parentNode = useNodeParent(rectRef);
const parentRect = useNodeClientRect(parentNode);

const config = computed((): Partial<Konva.RectConfig> => ({
  fill: 'rgba(255, 0, 0, 0.1)',
  x: 0,
  y: 0,
  width: parentRect.value.width,
  height: parentRect.value.height,
}));

defineExpose({
  debugNode: true,
});
</script>
