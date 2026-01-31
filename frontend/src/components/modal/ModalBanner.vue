<template>
  <Banner variant="control" size="sm">
    <slot />
  </Banner>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import { Banner } from '@/components/banner';

defineSlots<{
  default: Slot;
}>();
</script>
