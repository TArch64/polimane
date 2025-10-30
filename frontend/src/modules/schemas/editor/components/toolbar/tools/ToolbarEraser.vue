<template>
  <ToolbarButton
    title="Очистити"
    :active="toolsStore.isEraser"
    @click="activate"
  >
    <DropletOffIcon />
  </ToolbarButton>
</template>

<script setup lang="ts">
import { useCanvasStore, useToolsStore } from '@editor/stores';
import { useHotKeys } from '@editor/composables';
import { EditorCursor, EditorTool } from '@editor/enums';
import { DropletOffIcon } from '@/components/icon';
import ToolbarButton from '../ToolbarButton.vue';

const toolsStore = useToolsStore();
const canvasStore = useCanvasStore();

function activate() {
  toolsStore.activateTool(EditorTool.ERASER);
  canvasStore.setCursor(EditorCursor.CROSSHAIR);
}

useHotKeys({
  mac: { Meta_Digit0: activate },
  win: { Ctrl_Digit0: activate },
});
</script>
