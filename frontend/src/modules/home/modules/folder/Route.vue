<template>
  <p>
    {{ folderId }}
  </p>
</template>

<script setup lang="ts">
import { useHomeStore } from '@/modules/home/stores';
import { definePreload } from '@/router/define';
import { lazyDestroyStore } from '@/helpers';
import { useFolderStore } from './stores';

defineProps<{
  folderId: string;
}>();

defineOptions({
  beforeRouteEnter: definePreload<'home-folder'>((route) => {
    const { folderId } = route.params;
    return useFolderStore().load(folderId);
  }),

  beforeRouteLeave: async (_, __, next) => {
    lazyDestroyStore(useFolderStore);
    next();
  },
});

const homeStore = useHomeStore();
const folderStore = useFolderStore();

homeStore.setRouteConfig({
  title: folderStore.folder.name,
});
</script>
