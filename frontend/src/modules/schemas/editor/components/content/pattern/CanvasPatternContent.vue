<template>
  <CanvasStackV :config :gap="12" align="center">
    <CanvasRow
      v-for="row in pattern.content"
      :key="row.id"
      :row
    />
  </CanvasStackV>
</template>

<script setup lang="ts">
import Konva from 'konva';
import type { ISchemaPattern } from '@/models';
import { CanvasStackV } from '../base';
import { CanvasRow } from '../row';

defineProps<{
  pattern: ISchemaPattern;
}>();

const config: Partial<Konva.GroupConfig> = {
  x: -12,
};
</script>
