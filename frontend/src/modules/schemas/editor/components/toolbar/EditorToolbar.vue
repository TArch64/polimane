<template>
  <Card as="aside" class="editor-toolbar">
    <ToolbarGrid>
      <ToolbarEraser />
      <ToolbarPalette />
    </ToolbarGrid>
  </Card>
</template>

<script setup lang="ts">
import { Card } from '@/components/card';
import { ToolbarPalette } from './palette';
import ToolbarEraser from './ToolbarEraser.vue';
import ToolbarGrid from './ToolbarGrid.vue';
</script>

<style scoped>
@property --toolbar-button-size {
  syntax: "<length>";
  inherits: true;
  initial-value: 24px;
}

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
