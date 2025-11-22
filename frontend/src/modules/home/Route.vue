<template>
  <CommonLayout
    :title="homeStore.title"
    :selected="homeStore.selection?.ids.size"
    :selected-title="homeStore.selection?.title"
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
import { CommonLayout } from '@/components/layout';
import { useHomeStore } from './stores';
import { HomeTopBarActions } from './components';

defineProps<{
  folderId?: string; // used in child routes
}>();

const homeStore = useHomeStore();
</script>
