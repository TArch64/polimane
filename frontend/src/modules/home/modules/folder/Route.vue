<template>
  <FolderList v-if="schemasStore.schemas.length" />
  <FolderEmpty v-else />
</template>

<script setup lang="ts">
import { useHomeStore } from '@/modules/home/stores';
import { definePreload } from '@/router/define';
import { lazyDestroyStore } from '@/helpers';
import { useFolderSchemasStore, useFolderStore } from './stores';
import { FolderEmpty, FolderList } from './components';
import { useSchemasSelection } from './composables';

defineProps<{
  folderId: string;
}>();

defineOptions({
  beforeRouteEnter: definePreload<'home-folder'>(async (route) => {
    const { folderId } = route.params;
    const folderStore = useFolderStore();
    const schemasStore = useFolderSchemasStore();

    await Promise.all([
      folderStore.load(folderId),
      schemasStore.load(folderId),
    ]);
  }),

  beforeRouteLeave: async (_, __, next) => {
    lazyDestroyStore(useFolderStore);
    lazyDestroyStore(useFolderSchemasStore);
    next();
  },
});

const homeStore = useHomeStore();
const folderStore = useFolderStore();
const schemasStore = useFolderSchemasStore();

homeStore.setRouteConfig({
  title: folderStore.folder.name,
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
    doMany: schemasStore.deleteMany,
  },

  // addSchemaToFolder: {
  //   getFolders: () => foldersStore.folders,
  //   do: foldersStore.addSchemas,
  // },
});
</script>
