<template>
  <Button class="color-eraser">
    <svg
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path d="M12 2.5C8.5 6.5 6.5 9 6.5 12C6.5 15.5 9 18 12 18C15 18 17.5 15.5 17.5 12C17.5 9 15.5 6.5 12 2.5Z" />
      <line x1="5" y1="19" x2="19" y2="5" />
    </svg>
  </Button>
</template>

<script setup lang="ts">
import { Button } from '@/components/button';
</script>

<style scoped>
@layer page {
  .color-eraser {
    padding: 2px;
    width: var(--color-button-size);
    height: var(--color-button-size);
    --button-background: color-mix(in srgb, var(--button-base-color), transparent 95%)
  }

  .color-eraser svg {
    width: 100%;
    height: 100%;
  }
}
</style>
