<template>
  <Button
    icon
    :active
    :disabled
    size="none"
    variant="secondary"
    class="toolbar-button"
  >
    <slot />
  </Button>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import { Button } from '@/components/button';

withDefaults(defineProps<{
  active?: boolean;
  disabled?: boolean;
}>(), {
  active: false,
  disabled: false,
});

defineSlots<{
  default: Slot;
}>();
</script>

<style scoped>
@layer page {
  .toolbar-button {
    padding: 0;
    width: var(--toolbar-button-size);
    height: var(--toolbar-button-size);
    --button-foreground: var(--color-text-2);
  }
}
</style>
