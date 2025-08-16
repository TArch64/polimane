<template>
  <div class="app-loader">
    <LogoIcon size="48" />
  </div>
</template>

<script setup lang="ts">
import { LogoIcon } from '@/components/icon';
</script>

<style scoped>
@layer page {
  .app-loader {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 999999;
    background-color: var(--color-background-2);
    display: flex;
    align-items: center;
    justify-content: center;
  }
}
</style>
