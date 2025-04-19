<template>
  <div class="sidebar-structure-item">
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
import { type Slot } from 'vue';
import { MoreHorizontalIcon } from '@/components/icon';
import { Dropdown } from '@/components/dropdown';
import { Button } from '@/components/button';
import { mergeAnchorName } from '@/helpers';

withDefaults(defineProps<{
  title: string;
  moreActionsButtonStyle?: Record<string, string> | Record<string, string>[];
}>(), {
  moreActionsButtonStyle: () => ({}),
});

defineSlots<{
  actions: Slot;
}>();
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

    &:hover {
      background-color: var(--color-background-2);

      .sidebar-structure-item__more-actions {
        opacity: 1;
      }
    }
  }

  .sidebar-structure-item__more-actions {
    opacity: 0;
    transition: opacity 0.15s ease-out;
  }
}
</style>
