<template>
  <div inert class="editor-selection" />
</template>

<script setup lang="ts">
import { useSelectionColor } from '@editor/composables';
import { useSelectionStore } from '@editor/stores';

const selectionStore = useSelectionStore();
const selectionColor = useSelectionColor();
</script>

<style scoped>
@layer page {
  .editor-selection {
    position: fixed;
    border: var(--divider);
    background-color: color-mix(in srgb, v-bind("selectionColor"), transparent 80%);
    top: v-bind("selectionStore.area.y + 'px'");
    left: v-bind("selectionStore.area.x + 'px'");
    width: v-bind("selectionStore.area.width + 'px'");
    height: v-bind("selectionStore.area.height + 'px'");
    border-radius: var(--rounded-sm);
    will-change: width, height, top, left;
  }
}
</style>
