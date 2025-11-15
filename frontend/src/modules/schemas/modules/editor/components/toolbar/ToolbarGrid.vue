<template>
  <Component :is="as" class="toolbar-grid">
    <slot />
  </Component>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import type { ComponentAs, NumericString } from '@/types';

withDefaults(defineProps<{
  as?: ComponentAs;
  columns: NumericString | number;
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
    grid-template-columns: repeat(v-bind("columns"), var(--toolbar-button-size));
  }
}
</style>
