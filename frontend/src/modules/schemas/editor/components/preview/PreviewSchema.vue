<template>
  <svg xmlns="http://www.w3.org/2000/svg" :viewBox :width :height>
    <PreviewSector
      v-for="{ sector, grid } of sectors"
      :key="sector"
      :schema
      :grid="grid.value"
    />
  </svg>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { ISchema } from '@/models';
import { BEAD_SIZE, useBeadsGrid } from '../../composables';
import { PreviewSector } from './PreviewSector';

const props = defineProps<{
  schema: ISchema;
}>();

function calcSideLength(size1: number, size2: number): number {
  return (size1 + size2 + 2) * BEAD_SIZE;
}

const width = computed(() => calcSideLength(props.schema.size.left, props.schema.size.right));
const height = computed(() => calcSideLength(props.schema.size.top, props.schema.size.bottom));

const viewBox = computed(() => [
  0,
  0,
  width.value,
  height.value,
].join(' '));

const sectors = useBeadsGrid(() => props.schema, {
  filter: (coord) => coord in props.schema.beads,
});
</script>
