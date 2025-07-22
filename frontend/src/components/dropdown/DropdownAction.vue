<template>
  <Button :to :danger :disabled class="dropdown-action">
    <Component class="dropdown-action__icon" :is="icon" />

    {{ title }}
  </Button>
</template>

<script setup lang="ts">
import type { Component } from 'vue';
import type { RouteLocationRaw } from 'vue-router';
import { Button } from '@/components/button';

withDefaults(defineProps<{
  icon: Component;
  title: string;
  danger?: boolean;
  disabled?: boolean;
  to?: RouteLocationRaw;
}>(), {
  danger: false,
  disabled: false,
});
</script>

<style scoped>
@layer components {
  .dropdown-action {
    display: flex;
    align-items: center;
    justify-content: flex-start;
  }

  .dropdown-action__icon {
    flex-shrink: 0;
    margin-right: 8px;
  }
}
</style>
