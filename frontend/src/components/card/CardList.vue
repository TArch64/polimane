<template>
  <ul class="card-list">
    <slot />
  </ul>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';

defineSlots<{
  default: Slot;
}>();
</script>

<style scoped>
@layer components {
  .card-list {
    margin: 0;
    padding: 0;
    list-style-type: none;
  }
}
</style>
