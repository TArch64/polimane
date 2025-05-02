<template>
  <div class="sidebar-structure-item" :class="classes" @click.stop="onMouseOver">
    {{ title }}

    <Dropdown>
      <template #activator="{ activatorStyle, open }">
        <Button
          icon
          class="sidebar-structure-item__more-actions"
          :style="mergeAnchorName(activatorStyle, moreActionsButtonStyle)"
          @click="open"
        >
          <MoreHorizontalIcon />
        </Button>
      </template>

      <slot name="actions" />
    </Dropdown>
  </div>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';
import { MoreHorizontalIcon } from '@/components/icon';
import { Dropdown } from '@/components/dropdown';
import { Button } from '@/components/button';
import { mergeAnchorName } from '@/helpers';
import type { ISchemaObject } from '@/models';
import { useActiveObjectStore } from '@/modules/schemas/editor/stores';

const props = withDefaults(defineProps<{
  object: ISchemaObject;
  title: string;
  moreActionsButtonStyle?: Record<string, string> | Record<string, string>[];
}>(), {
  moreActionsButtonStyle: () => ({}),
});

defineSlots<{
  actions: Slot;
}>();

const activeObjectStore = useActiveObjectStore();
const isActive = activeObjectStore.useActiveObject(() => props.object);
const onMouseOver = () => activeObjectStore.activateObject(props.object);

const classes = computed(() => ({
  'sidebar-structure-item--active': isActive.value,
}));
</script>

<style scoped>
@layer page {
  .sidebar-structure-item {
    font-size: var(--font-sm);
    padding: 4px 8px 4px 12px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: var(--color-background-1);
    transition: background-color 0.15s ease-out;
    will-change: background-color;

    &:hover:not(.sidebar-structure-item--active) {
      background-color: var(--color-background-2);

      .sidebar-structure-item__more-actions {
        opacity: 1;
      }
    }
  }

  .sidebar-structure-item--active {
    background-color: var(--color-background-3);

    .sidebar-structure-item__more-actions {
      opacity: 1;
    }
  }

  .sidebar-structure-item__more-actions {
    opacity: 0;
    transition: opacity 0.15s ease-out;
  }
}
</style>
