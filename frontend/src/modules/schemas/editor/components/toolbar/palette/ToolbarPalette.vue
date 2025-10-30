<template>
  <ToolbarDropdown>
    <template #activator="{ open }">
      <ColorItem
        title="Колір Бісеру"
        :color="store.activeColor"
        @click="open"
      />
    </template>

    <ColorSelector />
  </ToolbarDropdown>
</template>

<script setup lang="ts">
import { useToolsStore } from '@editor/stores';
import ToolbarDropdown from '../ToolbarDropdown.vue';
import ColorItem from './ColorItem.vue';
import ColorSelector from './ColorSelector.vue';

const store = useToolsStore();
</script>
