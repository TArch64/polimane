<template>
  <ul class="sidebar-structure-list">
    <Sortable
      :list
      :gap="0"
      item-as="li"
      group="sidebar-pattern-list"
      direction="vertical"
      @move="$emit('move', $event)"
      v-slot="{ item }"
    >
      <slot :object="item" />
    </Sortable>
  </ul>
</template>

<script setup lang="ts" generic="O extends ISchemaObject">
import type { Slot } from 'vue';
import { type IMoveEvent, Sortable } from '@/components/sortable';
import type { ISchemaObject } from '@/models';

defineProps<{
  list: O[];
  sortableGroup: string;
}>();

defineEmits<{
  move: [event: IMoveEvent<O>];
}>();

defineSlots<{
  default: Slot<{ object: O }>;
}>();

</script>

<style scoped>
@layer page {
  .sidebar-structure-list {
    list-style-type: none;
    padding: 0;
    margin: 0;
  }
}
</style>
