<template>
  <ToolbarDropdown ref="dropdownRef">
    <template #activator>
      <ToolbarButton
        title="Бісер"
        :active="toolsStore.isBead"
        @click="onActivatorClick"
      >
        <BeadIcon :kind="toolsStore.activeBead" />
      </ToolbarButton>
    </template>

    <ToolbarGrid columns="1">
      <ToolbarButton
        v-for="kind of BeadContentList"
        :key="kind"
        :title="getBeadKindTitle(kind)"
        :active="toolsStore.activeBead === kind"
        @click="activateBead(kind)"
      >
        <BeadIcon :kind />
      </ToolbarButton>
    </ToolbarGrid>
  </ToolbarDropdown>
</template>

<script setup lang="ts">
import { useCanvasStore, useToolsStore } from '@editor/stores';
import { EditorCursor, EditorTool } from '@editor/enums';
import { ref } from 'vue';
import type { ComponentExposed } from 'vue-component-type-helpers';
import { type BeadContentKind, BeadContentList, getBeadKindTitle } from '@/enums';
import ToolbarButton from '../ToolbarButton.vue';
import ToolbarDropdown from '../ToolbarDropdown.vue';
import ToolbarGrid from '../ToolbarGrid.vue';
import { BeadIcon } from './BeadIcon';

const toolsStore = useToolsStore();
const canvasStore = useCanvasStore();

const dropdownRef = ref<ComponentExposed<typeof ToolbarDropdown>>(null!);

function activate() {
  toolsStore.activateTool(EditorTool.BEAD);
  canvasStore.setCursor(EditorCursor.CROSSHAIR);
}

function onActivatorClick(): void {
  toolsStore.isBead
    ? dropdownRef.value.open()
    : activate();
}

function activateBead(bead: BeadContentKind): void {
  toolsStore.activateBead(bead);
  dropdownRef.value.close();
}
</script>
