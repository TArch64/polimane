<template>
  <HomeList v-if="schemasStore.hasSchemas" />
  <HomeListEmpty v-else />
</template>

<script setup lang="ts">
import { definePreload } from '@/router/define';
import { useHomeStore } from '../../stores';
import { useFoldersStore, useHomeListStore, useSchemasStore } from './stores';
import { useSchemasSelection } from './composables';
import { HomeList, HomeListEmpty } from './components';

defineOptions({
  beforeRouteEnter: definePreload<'home'>(async () => {
    await useHomeListStore().load();
  }),
});

const homeStore = useHomeStore();
const schemasStore = useSchemasStore();
const foldersStore = useFoldersStore();

homeStore.setSelection(useSchemasSelection());

homeStore.setStrategies({
  createSchema: {
    do: schemasStore.createSchema,
  },

  addSchemaToFolder: {
    getFolders: () => foldersStore.folders,
    do: foldersStore.addSchemas,
  },
});
</script>
