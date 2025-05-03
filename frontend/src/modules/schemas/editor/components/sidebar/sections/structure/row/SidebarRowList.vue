<template>
  <SidebarStructureList
    :depth
    :list="pattern.content"
    sortable-group="sidebar-pattern-list"
    @move="onMove"
    @mouseout="hoverObjectStore.deactivatePath"
    v-slot="{ object }"
  >
    <SidebarRow :depth :pattern :row="object" />
  </SidebarStructureList>
</template>

<script setup lang="ts">
import { useHoverObjectStore, useRowsStore } from '@/modules/schemas/editor/stores';
import type { IMoveEvent } from '@/components/sortable';
import type { ISchemaPattern, ISchemaRow } from '@/models';
import { SidebarStructureList } from '../base';
import SidebarRow from './SidebarRow.vue';

const props = defineProps<{
  depth: number;
  pattern: ISchemaPattern;
}>();

const rowsStore = useRowsStore(() => props.pattern);
const hoverObjectStore = useHoverObjectStore();

function onMove(event: IMoveEvent<ISchemaRow>): void {
  rowsStore.rows.values = event.updated;
}
</script>
