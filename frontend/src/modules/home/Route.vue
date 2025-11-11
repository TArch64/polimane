<template>
  <CommonLayout
    :selected="schemasStore.selected.size"
    @clear-selection="schemasStore.clearSelection"
  >
    <template #top-bar-actions>
      <HomeTopBarActions />
    </template>

    <template #selection-title="{ count }">
      Обрано {{ count }} схем
    </template>

    <template #selection-actions>
      <HomeSelectionBarActions />
    </template>

    <HomeSchemasList v-if="schemasStore.hasSchemas" />
    <HomeSchemasEmpty v-else />
  </CommonLayout>
</template>

<script setup lang="ts">
import { definePreload } from '@/router/define';
import { CommonLayout } from '@/components/layout';
import {
  HomeSchemasEmpty,
  HomeSchemasList,
  HomeSelectionBarActions,
  HomeTopBarActions,
} from './components';
import { useSchemasStore } from './stores';

defineOptions({
  beforeRouteEnter: definePreload<'home'>(async () => {
    await useSchemasStore().load();
  }),
});

const schemasStore = useSchemasStore();
</script>
