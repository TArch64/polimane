<template>
  <SidebarStructureList
    :depth
    :list="beadsStore.beads.values"
    sortable-group="sidebar-pattern-list"
    @mouseout="hoverObjectStore.deactivatePath"
    v-slot="{ object }"
  >
    <SidebarBead :depth :row :bead="object" />
  </SidebarStructureList>
</template>

<script setup lang="ts">
import { useBeadsStore, useHoverObjectStore } from '@/modules/schemas/editor/stores';
import type { ISchemaRow } from '@/models';
import { SidebarStructureList } from '../base';
import SidebarBead from './SidebarBead.vue';

const props = defineProps<{
  depth: number;
  row: ISchemaRow;
}>();

const beadsStore = useBeadsStore(() => props.row);
const hoverObjectStore = useHoverObjectStore();
</script>
