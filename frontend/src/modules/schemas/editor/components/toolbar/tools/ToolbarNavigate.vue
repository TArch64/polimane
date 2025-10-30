<template>
  <ToolbarButton
    title="Переміщення Полотна"
    :active="toolsStore.isNavigate"
    @click="activate"
  >
    <MoveIcon />
  </ToolbarButton>
</template>

<script setup lang="ts">
import { useCanvasStore, useToolsStore } from '@editor/stores';
import { EditorCursor, EditorCursorTarget, EditorTool } from '@editor/enums';
import { useHotKeys } from '@editor/composables';
import { MoveIcon } from '@/components/icon';
import ToolbarButton from '../ToolbarButton.vue';

const toolsStore = useToolsStore();
const canvasStore = useCanvasStore();

function activate() {
  toolsStore.activateTool(EditorTool.NAVIGATE);
  canvasStore.setCursor(EditorCursor.GRAB, EditorCursorTarget.CANVAS);
}

useHotKeys({
  mac: { Alt_KeyH: activate },
  win: { Ctrl_KeyH: activate },
});
</script>
