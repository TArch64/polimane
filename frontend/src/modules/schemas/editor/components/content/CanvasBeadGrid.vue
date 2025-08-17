<template>
  <CanvasBead
    v-for="bead of beads"
    :key="bead.coord"
    :offset="bead.offset"
    :coord="bead.coord"
    :color="beadsStore.getColor(bead.coord)"
  />
</template>

<script lang="ts" setup>
import { computed } from 'vue';
import { useBeadsStore } from '../../stores';
import { CanvasBead } from './CanvasBead';
import type { BeadGridGenerator } from './useBeadsGrid';

const props = defineProps<{
  grid: BeadGridGenerator;
}>();

const beadsStore = useBeadsStore();
const beads = computed(() => Array.from(props.grid));
</script>
