<template>
  <header class="common-layout-top-bar">
    <LogoIcon size="28" />

    <h1 class="common-layout-top-bar__title">
      Polimane
    </h1>

    <div class="common-layout-top-bar__actions">
      <slot />
    </div>
  </header>
</template>

<script setup lang="ts">
import type { Slot } from 'vue';
import { LogoIcon } from '@/components/icon';

defineSlots<{
  default: Slot;
}>();
</script>

<style scoped>
@layer components {
  .common-layout-top-bar {
    position: sticky;
    top: 0;
    display: flex;
    align-items: center;
    padding: 12px 16px;
    min-height: 60px;
    background-color: var(--color-background-1);
    border-bottom: var(--divider);
  }

  .common-layout-top-bar__title {
    margin-left: 8px;
    font-size: 20px;
    font-weight: 500;
  }

  .common-layout-top-bar__actions {
    margin-left: auto;
    margin-right: 2px;
    display: flex;
    gap: 8px;
    align-items: center;
  }
}
</style>
