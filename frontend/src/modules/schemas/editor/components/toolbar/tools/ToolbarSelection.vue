<template>
  <ToolbarButton
    title="Виділення"
    :active="toolsStore.isSelection"
    @click="activate"
  >
    <ToolsSelectionIcon />
  </ToolbarButton>
</template>

<script setup lang="ts">
import { useCanvasStore, useToolsStore } from '@editor/stores';
import { EditorCursor, EditorTool } from '@editor/enums';
import { ToolsSelectionIcon } from '@/components/icon';
import ToolbarButton from '../ToolbarButton.vue';

const toolsStore = useToolsStore();
const canvasStore = useCanvasStore();

function activate() {
  toolsStore.activateTool(EditorTool.SELECTION);
  canvasStore.setCursor(EditorCursor.DEFAULT);
}
</script>
