<template>
  <section class="modal-banner">
    <slot />
  </section>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';

defineSlots<{
  default: Slot;
}>();
</script>

<style scoped>
@layer components {
  .modal-banner {
    background-color: var(--color-background-2);
    padding: 8px 12px;
    border-radius: var(--rounded-md);
    font-size: var(--font-sm);

    :deep(p:not(:first-child)) {
      margin-top: 4px;
    }
  }
}
</style>
