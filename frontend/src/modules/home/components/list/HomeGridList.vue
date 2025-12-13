<template>
  <div class="home-grid-list">
    <CursorSelection
      :list
      v-model="selected"
      v-slot="ctx"
      v-if="selectable"
    >
      <slot v-bind="ctx" />
    </CursorSelection>

    <template v-else>
      <template v-for="item in list" :key="item.id">
        <slot :item />
      </template>
    </template>
  </div>
</template>

<script setup lang="ts" generic="I extends { id: string }">
import type { Slot, VNodeRef } from 'vue';
import { CursorSelection } from '@/components/selection';

withDefaults(defineProps<{
  list: I[];
  selectable?: boolean;
}>(), {
  selectable: false,
});

const selected = defineModel<Set<string>>('selected', {
  required: false,
  default: () => new Set<string>(),
});

defineSlots<{
  default: Slot<{
    item: I;
    itemRef?: VNodeRef;
  }>;
}>();
</script>

<style scoped>
@layer page {
  .home-grid-list {
    display: grid;
    grid-template-columns: repeat(var(--list-columns), 1fr);
    align-content: start;
    align-items: end;
    gap: 20px;
    padding: 12px;
    --list-columns: 4;
  }

  @media (max-width: 992px) {
    .home-grid-list {
      --list-columns: 3;
    }
  }

  @media (max-width: 768px) {
    .home-grid-list {
      --list-columns: 2;
    }
  }
}
</style>
