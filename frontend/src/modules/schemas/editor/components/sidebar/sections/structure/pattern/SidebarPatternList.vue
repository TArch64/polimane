<template>
  <SidebarStructureList
    :list="patternsStore.patterns.values"
    sortable-group="sidebar-pattern-list"
    @move="onMove"
    @mouseout="hoverObjectStore.deactivatePath"
    v-slot="{ object }"
  >
    <SidebarPattern :pattern="object" />
  </SidebarStructureList>
</template>

<script setup lang="ts">
import { useHoverObjectStore, usePatternsStore } from '@/modules/schemas/editor/stores';
import type { IMoveEvent } from '@/components/sortable';
import type { ISchemaPattern } from '@/models';
import { SidebarStructureList } from '../base';
import SidebarPattern from './SidebarPattern.vue';

const hoverObjectStore = useHoverObjectStore();
const patternsStore = usePatternsStore();

function onMove(event: IMoveEvent<ISchemaPattern>): void {
  patternsStore.patterns.values = event.updated;
}
</script>
