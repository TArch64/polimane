<template>
  <Button icon size="none" class="color-item" :class="classes">
    <slot />
  </Button>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';
import { Button } from '@/components/button';

const props = defineProps<{
  color: string;
  active: boolean;
}>();

defineSlots<{
  default: Slot;
}>();

const classes = computed(() => {
  const modifier = props.color ? 'value' : 'empty';

  return [
    `color-item--${modifier}`,
    { 'color-item--active': props.active },
  ];
});
</script>

<style scoped>
@layer page {
  .color-item {
    aspect-ratio: 1;
    padding: 0;
    border: var(--divider);
    transition: background-color 0.15s ease-out, border-color 0.15s ease-out;
    will-change: background-color, border-color;
  }

  .color-item--active {
    outline: solid 1px var(--color-primary);
    outline-offset: 1px;
  }

  .color-item--value {
    --button-background: v-bind("color");
    --button-hover-background: color-mix(in srgb, v-bind("color"), transparent 20%);
  }

  .color-item--empty {
    background-image: url("@/assets/emptyColor.svg");
    background-size: 100% 100%;
  }
}
</style>
