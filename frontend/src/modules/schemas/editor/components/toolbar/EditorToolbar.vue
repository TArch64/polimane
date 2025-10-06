<template>
  <Card as="aside" class="editor-toolbar">
    <ToolbarEraser />
    <ToolbarTool />

    <div class="editor-toolbar__color">
      <ToolbarBackgroundColor class="editor-toolbar__color-background" />

      <Transition name="editor-toolbar__color-foreground-" :duration="150">
        <ToolbarPalette class="editor-toolbar__color-foreground" v-if="!store.isEraser" />
      </Transition>
    </div>
  </Card>
</template>

<script setup lang="ts">
import { useToolsStore } from '@editor/stores';
import { Card } from '@/components/card';
import { ToolbarPalette } from './palette';
import { ToolbarEraser, ToolbarTool } from './tools';
import ToolbarBackgroundColor from './ToolbarBackgroundColor.vue';

const store = useToolsStore();
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
    align-items: center;
    gap: 8px;
    padding: 8px 6px;
    --toolbar-button-size: 28px;
  }

  .editor-toolbar__color {
    position: relative;
    padding-top: 4px;
    padding-bottom: calc(var(--toolbar-button-size) / 2);
    padding-right: calc(var(--toolbar-button-size) / 2);
    --toolbar-button-size: 24px;
  }

  .editor-toolbar__color-background {
    transition: translate 150ms, scale 150ms;
    will-change: translate, scale;

    &:last-child,
    &:has(+ .editor-toolbar__color-foreground--leave-active) {
      translate: 25% 25%;
      scale: calc(30 / 24);
    }
  }

  .editor-toolbar__color-foreground {
    position: absolute;
    top: calc(var(--toolbar-button-size) - (var(--toolbar-button-size) / 2));
    left: calc(var(--toolbar-button-size) - (var(--toolbar-button-size) / 2));
  }

  .editor-toolbar__color-foreground--enter-from,
  .editor-toolbar__color-foreground--leave-to {
    opacity: 0;
    scale: 0.8;
  }

  .editor-toolbar__color-foreground--enter-active,
  .editor-toolbar__color-foreground--leave-active {
    transition: opacity 150ms, scale 150ms;
    will-change: opacity, scale;
  }
}
</style>
