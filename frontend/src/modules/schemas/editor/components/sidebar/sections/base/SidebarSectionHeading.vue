<template>
  <h1 class="sidebar-section__heading">
    <slot />
  </h1>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';

defineSlots<{
  default: Slot;
}>();
</script>

<style scoped>
@layer page {
  .sidebar-section__heading {
    font-size: var(--font-md);
    font-weight: 550;
  }
}
</style>
