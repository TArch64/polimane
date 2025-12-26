<template>
  <p v-if="schemasStore.schemas.length">
    schemas
  </p>

  <DeletedEmpty v-else />
</template>

<script setup lang="ts">
import { definePreload } from '@/router/define';
import { useHomeStore } from '../../stores';
import { useSchemasSelection } from './composables';
import { useDeletedSchemasStore } from './stores';
import { DeletedEmpty } from './components';

defineOptions({
  beforeRouteEnter: definePreload<'home-recently-deleted'>(async () => {
    await useDeletedSchemasStore().load();
  }),
});

const homeStore = useHomeStore();
const schemasStore = useDeletedSchemasStore();

homeStore.setRouteConfig({
  title: 'Нещодавно Видалені Схеми',
  selection: useSchemasSelection(),

  // implement later
  addSchemaToFolder: null!,
  copySchema: null!,
  createSchema: null!,
  deleteSchema: null!,
  updateFolder: null!,
  updateSchema: null!,
});
</script>
