<template>
  <Card as="aside" class="common-layout-submenu">
    <slot />
  </Card>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import { Card } from '../card';

defineSlots<{
  default: Slot;
}>();
</script>

<style scoped>
@layer components {
  .common-layout-submenu {
    display: flex;
    flex-direction: column;
    gap: 4px;
  }

  :deep(.button) {
    justify-content: flex-start;
  }
}
</style>
