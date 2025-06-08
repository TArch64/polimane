<template>
  <KonvaRect ref="rootRef" :config />
</template>

<script setup lang="ts">
import Konva from 'konva';
import type { ISchemaBead } from '@/models';
import { useNodeCursor, useNodeRef } from '@/modules/schemas/editor/composables';

defineProps<{
  bead: ISchemaBead;
}>();

const config: Partial<Konva.RectConfig> = {
  width: 14,
  height: 14,
  fill: 'rgba(0, 0, 0, 0.05)',
};

const rootRef = useNodeRef<Konva.Rect>();
useNodeCursor(rootRef, 'crosshair');
</script>
