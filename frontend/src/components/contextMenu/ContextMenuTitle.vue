<template>
  <h2 class="context-menu__title text-truncate">
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
  .context-menu__title {
    font-size: var(--font-sm);
    font-weight: 570;
    max-width: 100%;
  }
}
</style>
