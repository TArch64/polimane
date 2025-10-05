<template>
  <Card as="aside" class="editor-toolbar">
    <ToolbarDivider />
    <ToolbarPalette />
  </Card>
</template>

<script setup lang="ts">
import { Card } from '@/components/card';
import { ToolbarPalette } from './palette';
import ToolbarDivider from './ToolbarDivider.vue';
</script>

<style scoped>
@layer page {
  .editor-toolbar {
    position: fixed;
    top: 62px;
    left: 8px;
    z-index: 10;
    display: flex;
    flex-direction: column;
    gap: 8px;
    padding: 8px 6px;
  }
}
</style>
