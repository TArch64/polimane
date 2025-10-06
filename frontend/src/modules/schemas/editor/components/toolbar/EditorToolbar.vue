<template>
  <Card as="aside" class="editor-toolbar">
    <ToolbarGrid>
      <ToolbarEraser />
    </ToolbarGrid>

    <div class="editor-toolbar__color">
      <ToolbarBackgroundColor class="editor-toolbar__color-background" />
      <ToolbarPalette class="editor-toolbar__color-foreground" />
    </div>
  </Card>
</template>

<script setup lang="ts">
import { Card } from '@/components/card';
import { ToolbarPalette } from './palette';
import ToolbarEraser from './ToolbarEraser.vue';
import ToolbarBackgroundColor from './ToolbarBackgroundColor.vue';
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
    width: calc(var(--toolbar-button-size) * 2 + 12px);
    z-index: 10;
    display: flex;
    flex-direction: column;
    gap: 8px;
    padding: 8px 6px;
  }

  .editor-toolbar__color {
    position: relative;
    padding-bottom: calc(var(--toolbar-button-size) / 2);
    padding-right: calc(var(--toolbar-button-size) / 2);
    align-self: center;
  }

  .editor-toolbar__color-background {
    width: var(--toolbar-button-size);
    height: var(--toolbar-button-size);
  }

  :deep(.editor-toolbar__color-foreground) {
    position: absolute;
    top: calc(var(--toolbar-button-size) - (var(--toolbar-button-size) / 2));
    left: calc(var(--toolbar-button-size) - (var(--toolbar-button-size) / 2));
    width: var(--toolbar-button-size);
    height: var(--toolbar-button-size);
  }
}
</style>
