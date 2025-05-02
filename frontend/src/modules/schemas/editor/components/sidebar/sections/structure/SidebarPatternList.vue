<template>
  <ul class="sidebar-pattern-list" @mouseout="hoverObjectStore.deactivatePath">
    <Sortable
      item-as="li"
      group="sidebar-pattern-list"
      direction="vertical"
      :gap="0"
      :list="patternsStore.patterns.values"
      @move="onMove"
      v-slot="{ item }"
    >
      <SidebarStructurePattern :pattern="item" />
    </Sortable>
  </ul>
</template>

<script setup lang="ts">
import { useHoverObjectStore, usePatternsStore } from '@/modules/schemas/editor/stores';
import { type IMoveEvent, Sortable } from '@/components/sortable';
import type { ISchemaPattern } from '@/models';
import { SidebarStructurePattern } from './pattern';

const hoverObjectStore = useHoverObjectStore();
const patternsStore = usePatternsStore();

function onMove(event: IMoveEvent<ISchemaPattern>): void {
  patternsStore.patterns.values = event.updated;
}
</script>

<style scoped>
@layer page {
  .sidebar-pattern-list {
    list-style-type: none;
    padding: 0;
    margin: 0;
  }
}
</style>
