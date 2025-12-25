<template />

<script setup lang="ts">
import { useHomeStore } from '../../stores';
import { useSchemasSelection } from './composables';

const homeStore = useHomeStore();

homeStore.setRouteConfig({
  title: 'Нещодавно Видалені Схеми',
  selection: useSchemasSelection(),

  // implement later
  addSchemaToFolder: undefined!,
  copySchema: undefined!,
  createSchema: undefined!,
  deleteSchema: undefined!,
  updateFolder: undefined!,
  updateSchema: undefined!,
});
</script>
