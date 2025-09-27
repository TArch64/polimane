<template>
  <div class="labeled-content">
    <label class="labeled-content__label">
      <slot />

      <span>
      {{ label }}
    </span>
    </label>
  </div>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';

defineProps<{
  label: string;
}>();

defineSlots<{
  default: Slot;
}>();
</script>

<style scoped>
@layer components {
  .labeled-content {
    display: flex;
  }

  .labeled-content__label {
    display: flex;
    align-items: center;
    gap: 16px;
  }
}
</style>
