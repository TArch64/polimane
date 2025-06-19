<template>
  <div role="menu" popover="manual" class="dropdown">
    <slot />
  </div>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';

defineSlots<{
  default: Slot;
}>();
</script>

<style scoped>
@layer components {
  .dropdown {
    background-color: var(--color-background-2);
    border: var(--divider);
    border-radius: var(--rounded-md);
    box-shadow: var(--box-shadow);
    display: flex;
    flex-direction: column;
    gap: 4px;
    width: max-content;
    view-transition-name: dropdown;
  }
}
</style>
