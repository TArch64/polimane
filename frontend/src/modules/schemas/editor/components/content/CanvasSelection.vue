<template>
  <div inert class="editor-selection" />
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useEditorStore, useSelectionStore } from '@editor/stores';
import { useContrast } from '@editor/composables';

const editorStore = useEditorStore();
const selectionStore = useSelectionStore();

const contrast = useContrast(() => editorStore.schema.backgroundColor, '#000');
const selectionColor = computed(() => contrast.value > 4.5 ? 'var(--color-black)' : 'var(--color-white)');
</script>

<style scoped>
@layer page {
  .editor-selection {
    position: fixed;
    border: var(--divider);
    background-color: color-mix(in srgb, v-bind("selectionColor"), transparent 80%);
    top: v-bind("selectionStore.selection.y + 'px'");
    left: v-bind("selectionStore.selection.x + 'px'");
    width: v-bind("selectionStore.selection.width + 'px'");
    height: v-bind("selectionStore.selection.height + 'px'");
    will-change: width, height, top, left;
  }
}
</style>
