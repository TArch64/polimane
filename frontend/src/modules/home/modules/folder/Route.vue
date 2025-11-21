<template>
  <p>
    {{ folderId }}
  </p>

  <p>
    {{ schemasStore.schemas }}
  </p>
</template>

<script setup lang="ts">
import { useHomeStore } from '@/modules/home/stores';
import { definePreload } from '@/router/define';
import { lazyDestroyStore } from '@/helpers';
import { useFolderSchemasStore, useFolderStore } from './stores';

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
});
</script>
