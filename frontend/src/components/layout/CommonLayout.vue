<template>
  <CommonLayoutTopBar>
    <slot name="top-bar-actions" />
  </CommonLayoutTopBar>

  <slot />
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import { usePageClass } from '@/composables';
import CommonLayoutTopBar from './CommonLayoutTopBar.vue';

defineSlots<{
  'default': Slot;
  'top-bar-actions'?: Slot;
}>();

usePageClass('app--common-layout');
</script>

<style scoped>
@layer components {
  :global(.app--common-layout) {
    background-color: var(--color-background-2);
  }
}
</style>
