<template>
  <SidebarStructureItem
    :title
    :depth
    :object="bead"
  />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { ISchemaBead, ISchemaRow } from '@/models';
import { useBeadsStore } from '@/modules/schemas/editor/stores';
import { SidebarStructureItem } from '../base';

const props = defineProps<{
  depth: number;
  bead: ISchemaBead;
  row: ISchemaRow;
}>();

const beadsStore = useBeadsStore(() => props.row);

const index = computed(() => beadsStore.beads.indexOf(props.bead));
const title = computed(() => `Бісеринка #${index.value + 1}`);
</script>
