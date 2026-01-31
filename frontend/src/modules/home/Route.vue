<template>
  <CommonLayout
    :title="homeStore.title"
    :selected="selectionCount"
    :selected-title="selectionTitle"
    :selected-actions="homeStore.selection.actions"
    @clear-selection="homeStore.selection.onClear()"
  >
    <template #top-bar-actions>
      <HomeTopBarActions />
    </template>

    <RouterView />
  </CommonLayout>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { CommonLayout } from '@/components/layout';
import { SCHEMA_PLURAL, usePageClass, usePluralFormatter } from '@/composables';
import { useHomeStore } from './stores';
import { HomeTopBarActions } from './components';

defineProps<{
  folderId?: string; // used in child routes
}>();

const homeStore = useHomeStore();

const selectionCount = computed(() => homeStore.selection.ids.size);
const selectionSchemaPlural = usePluralFormatter(selectionCount, SCHEMA_PLURAL);
const selectionTitle = computed(() => `Обрано ${selectionCount.value} ${selectionSchemaPlural.value}`);

usePageClass('app-layout--home');
</script>

<style>
@layer page {
  .app-layout--home .common-layout__main {
    gap: 32px;
    padding-top: 32px;
  }
}
</style>
