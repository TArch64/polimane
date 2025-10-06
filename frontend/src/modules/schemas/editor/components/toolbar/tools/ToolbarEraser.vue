<template>
  <ToolbarButton
    title="Очитити"
    class="color-eraser"
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

<style scoped>
@layer page {
  .color-eraser {
    padding: 1px;
    --button-background: color-mix(in srgb, var(--button-base-color), transparent 90%);
  }
}
</style>
