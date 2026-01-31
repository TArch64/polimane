<template>
  <Card :variant as="div" role="menu" popover="manual" class="dropdown-menu">
    <slot />
  </Card>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import type { ComponentVariant } from '@/types';
import { Card } from '../card';

withDefaults(defineProps<{
  variant?: ComponentVariant;
  viewTransitionName?: string;
}>(), {
  variant: 'main',
  viewTransitionName: 'dropdown',
});

defineSlots<{
  default: Slot;
}>();
</script>

<style scoped>
@layer components {
  .dropdown-menu {
    display: flex;
    flex-direction: column;
    gap: 4px;
    width: max-content;
    view-transition-name: v-bind("viewTransitionName");
    --card-padding-top: 4px;
    --card-padding-bottom: 4px;
    --card-padding-left: 4px;
    --card-padding-right: 4px;
  }
}
</style>
