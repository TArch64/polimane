<template>
  <ToolbarButton
    title="Виділення"
    :active="store.isSelection"
    @click="store.activateTool('selection')"
  >
    <ToolsSelectionIcon />
  </ToolbarButton>
</template>

<script setup lang="ts">
import { useToolsStore } from '@editor/stores';
import { ToolsSelectionIcon } from '@/components/icon';
import ToolbarButton from '../ToolbarButton.vue';

const store = useToolsStore();
</script>
