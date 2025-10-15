<template>
  <Component :is="as" class="toolbar-grid">
    <slot />
  </Component>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import type { ComponentAs } from '@/types';

withDefaults(defineProps<{
  as?: ComponentAs;
}>(), {
  as: 'div',
});

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
