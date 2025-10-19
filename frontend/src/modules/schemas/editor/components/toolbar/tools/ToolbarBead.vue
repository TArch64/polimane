<template>
  <ToolbarDropdown ref="dropdownRef">
    <template #activator>
      <ToolbarButton
        title="Бісер"
        :active="store.isBead"
        @click="onActivatorClick"
      >
        <BeadIcon :kind="store.activeBead" />
      </ToolbarButton>
    </template>

    <ToolbarGrid columns="1">
      <ToolbarButton
        v-for="kind of BeadContentList"
        :key="kind"
        :title="getBeadKindTitle(kind)"
        @click="activateBead(kind)"
      >
        <BeadIcon :kind />
      </ToolbarButton>
    </ToolbarGrid>
  </ToolbarDropdown>
</template>

<script setup lang="ts">
import { useToolsStore } from '@editor/stores';
import { EditorTool } from '@editor/enums';
import { ref } from 'vue';
import type { ComponentExposed } from 'vue-component-type-helpers';
import { type BeadContentKind, BeadContentList, getBeadKindTitle } from '@/enums';
import ToolbarButton from '../ToolbarButton.vue';
import ToolbarDropdown from '../ToolbarDropdown.vue';
import ToolbarGrid from '../ToolbarGrid.vue';
import { BeadIcon } from './BeadIcon';

const store = useToolsStore();

const dropdownRef = ref<ComponentExposed<typeof ToolbarDropdown>>(null!);

function onActivatorClick(): void {
  store.isBead
    ? dropdownRef.value.open()
    : store.activateTool(EditorTool.BEAD);
}

function activateBead(bead: BeadContentKind): void {
  store.activateBead(bead);
  dropdownRef.value.close();
}
</script>
