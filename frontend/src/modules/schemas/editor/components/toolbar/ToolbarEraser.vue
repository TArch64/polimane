<template>
  <Button
    icon
    size="none"
    title="Очитити"
    class="color-eraser"
    :class="classes"
    @click="activate"
  >
    <svg
      class="color-eraser__icon"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M12 2.5C8.5 6.5 6.5 9 6.5 12C6.5 15.5 9 18 12 18C15 18 17.5 15.5 17.5 12C17.5 9 15.5 6.5 12 2.5Z"
      />

      <line
        x1="5"
        y1="19"
        x2="19"
        y2="5"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
  </Button>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useToolsStore } from '@editor/stores';
import { Button } from '@/components/button';

const store = useToolsStore();
const activate = () => store.activateTool('eraser');

const classes = computed(() => ({
  'color-eraser--active': store.activeToolId === 'eraser',
}));
</script>

<style scoped>
@layer page {
  .color-eraser {
    padding: 2px;
    aspect-ratio: 1;
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
