<template>
  <footer class="editor-bottom-bar">
    <BottomBarBeadsCount class="editor-bottom-bar__divider-left editor-bottom-bar__pull-right" />
    <BottomBarZoom class="editor-bottom-bar__divider-left" />
  </footer>
</template>

<script setup lang="ts">
import { BottomBarBeadsCount } from './beadsCount';
import BottomBarZoom from './BottomBarZoom.vue';
</script>

<style scoped>
@layer page {
  .editor-bottom-bar {
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
    z-index: 10;
    background-color: var(--color-background-1);
    font-size: var(--font-sm);
    color: var(--color-text-2);
    border-top: var(--divider);
    display: flex;
    height: 20px;
  }

  .editor-bottom-bar__pull-right {
    margin-left: auto;
  }

  .editor-bottom-bar__divider-left {
    border-left: var(--divider);
  }
}
</style>
