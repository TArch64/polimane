<template>
  <ToolbarButton
    title="Збільшення/Зменшення"
    :active="toolsStore.isZoom"
    @click="activate"
  >
    <MaximizeIcon />
  </ToolbarButton>
</template>

<script setup lang="ts">
import { useCanvasStore, useToolsStore } from '@editor/stores';
import { EditorCursor, EditorCursorTarget, EditorTool } from '@editor/enums';
import { useHotKeys } from '@editor/composables';
import { MaximizeIcon } from '@/components/icon';
import ToolbarButton from '../ToolbarButton.vue';

const toolsStore = useToolsStore();
const canvasStore = useCanvasStore();

function activate() {
  toolsStore.activateTool(EditorTool.ZOOM);
  canvasStore.setCursor(EditorCursor.ZOOM_IN, EditorCursorTarget.CANVAS);
}

useHotKeys({
  Alt_KeyH: activate,
});
</script>
