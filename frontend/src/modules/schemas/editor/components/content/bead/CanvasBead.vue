<template>
  <KonvaRect :config />
</template>

<script setup lang="ts">
import Konva from 'konva';
import type { ISchemaBead } from '@/models';

defineProps<{
  bead: ISchemaBead;
}>();

const config: Partial<Konva.RectConfig> = {
  width: 14,
  height: 14,
  fill: 'rgba(0, 0, 0, 0.05)',
};
</script>
