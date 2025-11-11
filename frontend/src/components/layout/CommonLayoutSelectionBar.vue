<template>
  <header class="common-layout-selection-bar">
    <h2 class="common-layout-selection-bar__title">
      <slot name="title" :count="selected" />
    </h2>

    <div class="common-layout-selection-bar__actions">
      <slot name="actions" />

      <Button
        icon
        size="md"
        variant="secondary"
        @click="$emit('clear-selection')"
      >
        <CloseIcon />
      </Button>
    </div>
  </header>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import { Button } from '../button';
import { CloseIcon } from '../icon';

defineProps<{
  selected: number;
}>();

defineEmits<{
  'clear-selection': [];
}>();

defineSlots<{
  title: Slot<{ count: number }>;
  actions: Slot;
}>();
</script>

<style scoped>
@layer components {
  .common-layout-selection-bar {
    position: sticky;
    top: 0;
    display: flex;
    align-items: center;
    padding: 12px 16px;
    min-height: 60px;
    background-color: var(--color-primary);
    border-bottom: var(--divider);
    color: var(--color-white);
    z-index: 10;
    --button-base-color: var(--color-white);
  }

  .common-layout-selection-bar__title {
    font-size: var(--font-md);
    font-weight: 450;
  }

  .common-layout-selection-bar__actions {
    margin-left: auto;
    margin-right: 2px;
    display: flex;
    gap: 8px;
    align-items: center;
    font-size: var(--font-sm);
  }
}
</style>
