<template>
  <section class="sidebar__section">
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
@layer page {
  .sidebar__section {
    padding: 16px 12px;

    &:not(:last-of-type) {
      border-bottom: var(--divider);
    }
  }
}
</style>
