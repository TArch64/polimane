<template>
  <ToolbarGrid columns="2">
    <ColorSelectorSpot
      v-for="(color, index) of store.palette"
      :key="index"
      :active="color === store.activeColor"
      :model-value="color"
      @update:active="store.activateColor(color)"
      @update:model-value="updateColor(index, $event)"
    />
  </ToolbarGrid>
</template>

<script setup lang="ts">
import { useToolsStore } from '@editor/stores';
import ToolbarGrid from '../ToolbarGrid.vue';
import ColorSelectorSpot from './ColorSelectorSpot.vue';

const store = useToolsStore();

function updateColor(index: number, color: string): void {
  const isActive = store.activeColor === store.palette[index];
  store.palette[index] = color;
  if (isActive) store.activateColor(color);
}
</script>
