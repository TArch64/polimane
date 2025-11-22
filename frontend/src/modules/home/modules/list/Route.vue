<template>
  <HomeList v-if="schemasStore.hasSchemas" />

  <HomeListEmpty
    description="Поки що не створено жодної схеми для бісеру"
    v-else
  />
</template>

<script setup lang="ts">
import { definePreload } from '@/router/define';
import { HomeListEmpty } from '@/modules/home/components';
import { useHomeStore } from '../../stores';
import { useFoldersStore, useHomeListStore, useSchemasStore } from './stores';
import { useSchemasSelection } from './composables';
import { HomeList } from './components';

defineOptions({
  beforeRouteEnter: definePreload<'home'>(async () => {
    await useHomeListStore().load();
  }),
});

const homeStore = useHomeStore();
const schemasStore = useSchemasStore();
const foldersStore = useFoldersStore();

homeStore.setRouteConfig({
  selection: useSchemasSelection(),

  createSchema: {
    do: schemasStore.createSchema,
  },

  updateSchema: {
    do: schemasStore.updateSchema,
  },

  copySchema: {
    do: schemasStore.copySchema,
  },

  deleteSchema: {
    do: schemasStore.deleteSchema,
  },

  addSchemaToFolder: {
    getFolders: () => foldersStore.folders,
    do: foldersStore.addSchemas,
  },
});
</script>
