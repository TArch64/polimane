<template>
  <ToolbarButton
    class="toolbar-tool"
    :active="store.isBead"
    @click="store.activateTool('bead')"
  >
    <ToolsBeadIcon />
  </ToolbarButton>
</template>

<script setup lang="ts">
import { useToolsStore } from '@editor/stores';
import { ToolsBeadIcon } from '@/components/icon';
import ToolbarButton from '../ToolbarButton.vue';

const store = useToolsStore();
</script>
