<template>
  <CommonLayout
    :selected="homeStore.selection?.count"
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

const homeStore = useHomeStore();
</script>
