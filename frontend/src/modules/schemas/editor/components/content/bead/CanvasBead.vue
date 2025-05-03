<template>
  <GroupRenderer>
    <KonvaRect :config="outerConfig" />
    <KonvaRect :config="innerConfig" />
  </GroupRenderer>
</template>

<script setup lang="ts">
import Konva from 'konva';
import type { ISchemaBead } from '@/models';
import { GroupRenderer } from '../base';

defineProps<{
  bead: ISchemaBead;
}>();

const SIZE = 16;

const outerConfig: Partial<Konva.GroupConfig> = {
  width: SIZE,
  height: SIZE,
};

const innerConfig: Partial<Konva.RectConfig> = {
  width: SIZE - 2,
  height: SIZE - 2,
  fill: 'rgba(0, 0, 0, 0.05)',
};
</script>
