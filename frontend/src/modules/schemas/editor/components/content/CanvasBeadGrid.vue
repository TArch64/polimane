<template>
  <CanvasBead
    v-for="bead of grid"
    :key="bead.coord"
    :offset="bead.offset"
    :coord="bead.coord"
    :color="beadsStore.getColor(bead.coord)"
  />
</template>

<script lang="ts" setup>
import { useBeadsStore } from '../../stores';
import { CanvasBead } from './CanvasBead';
import type { IBeadsGridItem } from './useBeadsGrid';

defineProps<{
  grid: IBeadsGridItem[];
}>();

const beadsStore = useBeadsStore();
</script>
