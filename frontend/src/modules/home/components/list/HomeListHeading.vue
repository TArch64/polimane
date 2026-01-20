<template>
  <h2 class="home-list-heading">
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
@layer page {
  .home-list-heading {
    padding-left: 16px;
    margin-bottom: 16px;
    font-weight: 450;
    font-size: var(--font-lg);
    color: var(--color-text-2);
    display: flex;
    align-items: center;
  }
}
</style>
