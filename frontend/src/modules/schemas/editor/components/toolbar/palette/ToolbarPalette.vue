<template>
  <ToolbarDropdown :offset-top="-16">
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
import { type HotKeyDef, useHotKeys } from '@editor/composables';
import ToolbarDropdown from '../ToolbarDropdown.vue';
import ColorItem from './ColorItem.vue';
import ColorSelector from './ColorSelector.vue';

const store = useToolsStore();

useHotKeys(
  store.palette.slice(0, 9).map((_, index): HotKeyDef => [
    `Meta_${index + 1}`,
    () => store.activateColor(store.palette[index]!),
  ]),
);
</script>
