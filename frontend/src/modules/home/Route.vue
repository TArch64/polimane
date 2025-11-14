<template>
  <CommonLayout
    :selected="selection.count"
    :selected-title="selection.title"
    :selected-actions="selection.actions"
    @clear-selection="schemasStore.clearSelection"
  >
    <template #top-bar-actions>
      <HomeTopBarActions />
    </template>

    <HomeSchemasList v-if="schemasStore.hasSchemas" />
    <HomeSchemasEmpty v-else />
  </CommonLayout>
</template>

<script setup lang="ts">
import { definePreload } from '@/router/define';
import { CommonLayout } from '@/components/layout';
import { HomeSchemasEmpty, HomeSchemasList, HomeTopBarActions } from './components';
import { useHomeListStore, useSchemasStore } from './stores';
import { useSchemasSelection } from './composables';

defineOptions({
  beforeRouteEnter: definePreload<'home'>(async () => {
    await useHomeListStore().load();
  }),
});

const schemasStore = useSchemasStore();
const selection = useSchemasSelection();
</script>
