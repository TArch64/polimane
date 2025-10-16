<template>
  <ToolbarButton
    title="Очистити"
    :active="store.isEraser"
    @click="activate"
  >
    <DropletOffIcon />
  </ToolbarButton>
</template>

<script setup lang="ts">
import { useToolsStore } from '@editor/stores';
import { useHotKeys } from '@editor/composables';
import { EditorTool } from '@editor/enums';
import { DropletOffIcon } from '@/components/icon';
import ToolbarButton from '../ToolbarButton.vue';

const store = useToolsStore();
const activate = () => store.activateTool(EditorTool.ERASER);

useHotKeys({
  Meta_0: activate,
});
</script>
