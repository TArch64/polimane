<template>
  <Button
    icon
    size="none"
    title="Очитити"
    class="color-eraser"
    :class="classes"
    @click="activate"
  >
    <CloseIcon />
  </Button>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useToolsStore } from '@editor/stores';
import { useHotKeys } from '@editor/composables';
import { Button } from '@/components/button';
import { CloseIcon } from '@/components/icon';

const store = useToolsStore();
const activate = () => store.activateTool('eraser');

const classes = computed(() => ({
  'color-eraser--active': store.activeToolId === 'eraser',
}));

useHotKeys({
  Meta_0: activate,
});
</script>

<style scoped>
@layer page {
  .color-eraser {
    aspect-ratio: 1;
    border: var(--divider);
    padding: 1px;
    --button-background: color-mix(in srgb, var(--button-base-color), transparent 90%);
  }

  .color-eraser--active {
    outline: var(--color-primary) solid 1px;
    outline-offset: 1px;
  }

  .color-eraser__icon {
    width: 100%;
    height: 100%;
  }
}
</style>
