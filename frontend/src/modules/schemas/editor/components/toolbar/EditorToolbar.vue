<template>
  <Card as="aside" class="editor-toolbar">
    <ToolbarPalette />
  </Card>
</template>

<script setup lang="ts">
import { Card } from '@/components/card';
import { ToolbarPalette } from './palette';
</script>

<style scoped>
@layer page {
  .editor-toolbar {
    position: fixed;
    top: 62px;
    left: 8px;
    z-index: 10;
    display: flex;
    padding: 8px 6px;
  }
}
</style>
