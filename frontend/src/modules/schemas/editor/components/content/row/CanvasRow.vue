<template>
  <KonvaRect :config />
</template>

<script setup lang="ts">
import Konva from 'konva';
import type { ISchemaRow } from '@/models';

defineProps<{
  row: ISchemaRow;
}>();

const config: Partial<Konva.RectConfig> = {
  width: 800,
  height: 20,
  fill: 'rgba(0, 0, 0, 0.2)',
};
</script>
