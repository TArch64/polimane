<template>
  <div class="sidebar-structure-item" :class="classes">
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

const props = withDefaults(defineProps<{
  title: string;
  active: boolean;
  moreActionsButtonStyle?: Record<string, string> | Record<string, string>[];
}>(), {
  moreActionsButtonStyle: () => ({}),
});

defineSlots<{
  actions: Slot;
}>();

const classes = computed(() => ({
  'sidebar-structure-item--active': props.active,
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
    }

    &:hover,
    &.sidebar-structure-item--active {

      .sidebar-structure-item__more-actions {
        opacity: 1;
      }
    }
  }

  .sidebar-structure-item--active {
    background-color: var(--color-background-3);
  }

  .sidebar-structure-item__more-actions {
    opacity: 0;
    transition: opacity 0.15s ease-out;
  }
}
</style>
