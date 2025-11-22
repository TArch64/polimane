<template>
  <CommonLayout
    :title="homeStore.title"
    :selected="selectionCount"
    :selected-title="selectionTitle"
    :selected-actions="homeStore.selection?.actions"
    @clear-selection="homeStore.selection?.onClear()"
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
import { useHomeStore } from './stores';
import { HomeTopBarActions } from './components';

defineProps<{
  folderId?: string; // used in child routes
}>();

const homeStore = useHomeStore();

const selectionCount = computed(() => homeStore.selection?.ids.size ?? 0);
const selectionTitle = computed(() => `Обрано ${selectionCount.value} схем`);
</script>
