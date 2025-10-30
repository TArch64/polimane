<template>
  <ToolbarButton
    :title="tooltip"
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
import { computed } from 'vue';
import { MoveIcon } from '@/components/icon';
import ToolbarButton from '../ToolbarButton.vue';

const toolsStore = useToolsStore();
const canvasStore = useCanvasStore();

function activate() {
  toolsStore.activateTool(EditorTool.NAVIGATE);
  canvasStore.setCursor(EditorCursor.GRAB, EditorCursorTarget.CANVAS);
}

const hotKeys = useHotKeys({
  mac: { Alt_KeyH: activate },
  win: { Ctrl_KeyH: activate },
});

const tooltip = computed(() => {
  const hotKey = hotKeys.titles.Alt_KeyH || hotKeys.titles.Ctrl_KeyH;
  return `Переміщення Полотна (${hotKey})`;
});
</script>
