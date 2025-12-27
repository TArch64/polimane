<template>
  <DeletedList v-if="schemasStore.schemas.length" />
  <DeletedEmpty v-else />
</template>

<script setup lang="ts">
import { definePreload } from '@/router/define';
import { lazyDestroyStore } from '@/helpers';
import { useHomeStore } from '../../stores';
import { useSchemasSelection } from './composables';
import { useDeletedSchemasStore } from './stores';
import { DeletedEmpty, DeletedList } from './components';

defineOptions({
  beforeRouteEnter: definePreload<'home-recently-deleted'>(async () => {
    await useDeletedSchemasStore().load();
  }),

  beforeRouteLeave: async (_, __, next) => {
    lazyDestroyStore(useDeletedSchemasStore);
    next();
  },
});

const homeStore = useHomeStore();
const schemasStore = useDeletedSchemasStore();

homeStore.setRouteConfig({
  title: 'Нещодавно Видалені Схеми',
  selection: useSchemasSelection(),

  // unavailable in current route
  addSchemaToFolder: null!,
  copySchema: null!,
  createSchema: null!,
  deleteSchema: null!,
  updateFolder: null!,
  updateSchema: null!,
});
</script>
