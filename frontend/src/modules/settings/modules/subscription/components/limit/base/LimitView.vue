<template>
  <div class="subscription__limit">
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
  .subscription__limit {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }
}
</style>
