<template>
  <p class="dropdown__text">
    <slot />
  </p>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';

defineSlots<{
  default: Slot;
}>();
</script>

<style scoped>
@layer components {
  .dropdown__text {
    display: flex;
    align-items: center;
    justify-content: flex-start;
    font-size: var(--font-sm);
    min-height: 30px;
    padding: 7px 12px;
    line-height: 18px;
  }
}
</style>
