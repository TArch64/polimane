<template>
  <div class="card">
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
  .card {
    padding: 8px;
    background-color: var(--color-background-2);
    border: 1px solid var(--color-divider);
    border-radius: var(--rounded-md);
  }
}
</style>
