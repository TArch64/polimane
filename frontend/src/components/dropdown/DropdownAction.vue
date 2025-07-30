<template>
  <Button
    :to
    :danger
    :disabled
    :prepend-icon="icon"
    class="dropdown-action"
  >
    {{ title }}
  </Button>
</template>

<script setup lang="ts">
import type { RouteLocationRaw } from 'vue-router';
import { Button } from '@/components/button';
import type { IconComponent } from '@/components/icon';

withDefaults(defineProps<{
  icon: IconComponent;
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
}
</style>
