<template>
  <div class="toolbar-grid">
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
@layer page {
  .toolbar-grid {
    display: grid;
    gap: 6px;
    grid-template-columns: repeat(2, var(--toolbar-button-size));
  }
}
</style>
