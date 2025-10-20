<template>
  <Card as="aside" ref="toolbarRef" class="editor-toolbar">
    <ToolbarEraser />
    <ToolbarBead />
    <ToolbarSelection />

    <div class="editor-toolbar__color">
      <div class="editor-toolbar__color-background">
        <ToolbarBackgroundColor />
      </div>

      <Transition name="editor-toolbar__color-foreground-" :duration="150">
        <ToolbarPalette class="editor-toolbar__color-foreground" v-if="store.isBead" />
      </Transition>
    </div>
  </Card>
</template>

<script setup lang="ts">
import { useToolsStore } from '@editor/stores';
import { useElementBounding } from '@vueuse/core';
import { computed } from 'vue';
import { Card } from '@/components/card';
import { useDomRef } from '@/composables';
import { ToolbarPalette } from './palette';
import { ToolbarBead, ToolbarEraser, ToolbarSelection } from './tools';
import ToolbarBackgroundColor from './ToolbarBackgroundColor.vue';
import { provideToolbarRef } from './toolbarRef';

const store = useToolsStore();

const toolbarRef = useDomRef<HTMLElement>();
provideToolbarRef(toolbarRef);

const toolbarSize = useElementBounding(toolbarRef, {
  windowScroll: false,
});

const toolbarTopShift = computed(() => `${toolbarSize.height.value / 2}px`);
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
    /* can't use translate since it breaks popover position */
    top: calc(50% - v-bind("toolbarTopShift"));
    left: var(--editor-ui-padding);
    z-index: 10;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 6px;
    padding: 8px 6px;
    --toolbar-button-size: 28px;
  }

  .editor-toolbar__color {
    position: relative;
    margin-top: 4px;
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
