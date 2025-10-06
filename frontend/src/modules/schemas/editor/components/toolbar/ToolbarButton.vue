<template>
  <Button
    icon
    :disabled
    size="none"
    class="toolbar-button"
    :class="classes"
  >
    <slot />
  </Button>
</template>

<script setup lang="ts">
import { computed, type Slot } from 'vue';
import { Button } from '@/components/button';

const props = withDefaults(defineProps<{
  active?: boolean;
  disabled?: boolean;
}>(), {
  active: false,
  disabled: false,
});

defineSlots<{
  default: Slot;
}>();

const classes = computed(() => ({
  'toolbar-button--active': props.active,
}));
</script>

<style scoped>
@layer page {
  .toolbar-button {
    width: var(--toolbar-button-size);
    height: var(--toolbar-button-size);
    padding: 0;
    border: var(--divider);
    color: var(--color-primary);
    transition: background-color 0.15s ease-out, border-color 0.15s ease-out;
    will-change: background-color, border-color;
  }

  .toolbar-button--active {
    outline: solid 1px var(--color-primary);
    outline-offset: 1px;
  }
}
</style>
