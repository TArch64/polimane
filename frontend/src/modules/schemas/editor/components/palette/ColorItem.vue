<template>
  <Button class="color-list__item" :class="classes" />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { Button } from '@/components/button';

const props = defineProps<{
  color: string;
}>();

const classes = computed(() => {
  const modifier = props.color ? 'with' : 'without';
  return `color-list__item--${modifier}-color`;
});
</script>

<style scoped>
@layer page {
  .color-list__item {
    padding: 0;
    border: var(--divider);
  }

  .color-list__item--with-color {
    background-color: v-bind("color");
  }

  .color-list__item--without-color {
    background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='100%25' height='100%25'%3E%3Crect width='100%25' height='100%25' fill='%23f8f8f8'/%3E%3Cline x1='0' y1='0' x2='100%25' y2='100%25' stroke='%23dddddd' stroke-width='1.5'/%3E%3C/svg%3E");
    background-size: cover;
  }
}
</style>
