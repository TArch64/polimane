<template>
  <h2 class="dropdown__title">
    <slot />
  </h2>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';

defineSlots<{
  default: Slot;
}>();
</script>

<style scoped>
@layer components {
  .dropdown__title {
    font-size: var(--font-sm);
    font-weight: 570;
  }
}
</style>
