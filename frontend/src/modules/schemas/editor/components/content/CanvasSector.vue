<template>
  <CanvasBead
    v-for="bead of grid"
    :key="bead.coord"
    :offset="bead.offset"
    :coord="bead.coord"
    :color="beadsStore.getColor(bead.coord)"
    :emptyColor="beadsStore.emptyColor"
  />
</template>

<script lang="ts" setup>
import { useBeadsStore } from '../../stores';
import type { IBeadsGridItem } from '../../composables';
import { CanvasBead } from './CanvasBead';

defineProps<{
  grid: IBeadsGridItem[];
}>();

const beadsStore = useBeadsStore();
</script>
