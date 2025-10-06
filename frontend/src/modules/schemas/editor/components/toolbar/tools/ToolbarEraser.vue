<template>
  <ToolbarButton
    title="Очитити"
    :active="store.isEraser"
    @click="activate"
  >
    <CloseIcon />
  </ToolbarButton>
</template>

<script setup lang="ts">
import { useToolsStore } from '@editor/stores';
import { useHotKeys } from '@editor/composables';
import { CloseIcon } from '@/components/icon';
import ToolbarButton from '../ToolbarButton.vue';

const store = useToolsStore();
const activate = () => store.activateTool('eraser');

useHotKeys({
  Meta_0: activate,
});
</script>
