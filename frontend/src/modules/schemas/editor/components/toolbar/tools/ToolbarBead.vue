<template>
  <ToolbarDropdown :offset-top="-8">
    <template #activator="{ open }">
      <ToolbarButton
        title="Бісер"
        :active="store.isBead"
        @click="store.activateTool(EditorTool.BEAD)"
        @dblclick="open"
      >
        <BeadIcon :kind="store.activeBead" />
      </ToolbarButton>
    </template>

    <template #default="{ close }">
      <ToolbarGrid columns="1">
        <ToolbarButton
          v-for="kind of BeadKindList"
          :key="kind"
          :title="getBeadKindTitle(kind)"
          @click="close(); store.activateBead(kind)"
        >
          <BeadIcon :kind />
        </ToolbarButton>
      </ToolbarGrid>
    </template>
  </ToolbarDropdown>
</template>

<script setup lang="ts">
import { useToolsStore } from '@editor/stores';
import { EditorTool } from '@editor/enums';
import { BeadKindList, getBeadKindTitle } from '@/enums';
import ToolbarButton from '../ToolbarButton.vue';
import ToolbarDropdown from '../ToolbarDropdown.vue';
import ToolbarGrid from '../ToolbarGrid.vue';
import { BeadIcon } from './BeadIcon';

const store = useToolsStore();
</script>
