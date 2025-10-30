<template>
  <ToolbarButton
    :title="tooltip"
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
import { computed } from 'vue';
import { DropletOffIcon } from '@/components/icon';
import ToolbarButton from '../ToolbarButton.vue';

const toolsStore = useToolsStore();
const canvasStore = useCanvasStore();

function activate() {
  toolsStore.activateTool(EditorTool.ERASER);
  canvasStore.setCursor(EditorCursor.CROSSHAIR);
}

const hotKeys = useHotKeys({
  mac: { Meta_Digit0: activate },
  win: { Ctrl_Digit0: activate },
});

const tooltip = computed(() => {
  const hotKey = hotKeys.titles.Meta_Digit0 || hotKeys.titles.Ctrl_Digit0;
  return `Очистити (${hotKey})`;
});
</script>
