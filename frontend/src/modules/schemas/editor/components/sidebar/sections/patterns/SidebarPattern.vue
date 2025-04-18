<template>
  <li class="sidebar-pattern" :class="classes">
    {{ pattern.name }}

    <Dropdown>
      <template #activator="{ activatorRef, open }">
        <Button
          icon
          class="sidebar-pattern__more-actions"
          :ref="activatorRef"
          @click="open"
        >
          <MoreHorizontalIcon />
        </Button>
      </template>

      <DropdownAction
        title="Переназвати Паттерн"
        :icon="EditIcon"
      />

      <DropdownAction
        danger
        title="Видалити Паттерн"
        :icon="TrashIcon"
        @click="deletePattern"
      />
    </Dropdown>
  </li>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { ISchemaPattern } from '@/models';
import { useEditorStore, usePatternsStore } from '@/modules/schemas/editor/stores';
import { Button } from '@/components/button';
import { Dropdown, DropdownAction } from '@/components/dropdown';
import { EditIcon, MoreHorizontalIcon, TrashIcon } from '@/components/icon';

const props = defineProps<{
  pattern: ISchemaPattern;
}>();

const editorStore = useEditorStore();
const patternsStore = usePatternsStore();

const classes = computed(() => ({
  'sidebar-pattern--active': editorStore.activePattern?.id === props.pattern.id,
}));

function deletePattern(): void {
  patternsStore.deletePattern(props.pattern);
}
</script>

<style scoped>
@layer page {
  .sidebar-pattern {
    font-size: var(--font-sm);
    padding: 4px 8px 4px 12px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    transition: background-color 0.15s ease-out;
    will-change: background-color;

    &:hover:not(.sidebar-pattern--active) {
      background-color: var(--color-background-2);
    }

    &:hover,
    &.sidebar-pattern--active {

      .sidebar-pattern__more-actions {
        opacity: 1;
      }
    }
  }

  .sidebar-pattern--active {
    background-color: var(--color-background-3);
  }

  .sidebar-pattern__more-actions {
    opacity: 0;
    transition: opacity 0.15s ease-out;
  }
}
</style>
