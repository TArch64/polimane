<template>
  <CommonLayout
    :selected="homeStore.selection?.count"
    :selected-title="homeStore.selection?.title"
    :selected-actions="homeStore.selection?.actions"
    @clear-selection="homeStore.selection?.onClear()"
  >
    <template #top-bar-actions>
      <div class="home-top-bar__actions" data-home-top-bar-actions />
    </template>

    <RouterView />
  </CommonLayout>
</template>

<script setup lang="ts">
import { CommonLayout } from '@/components/layout';
import { useHomeStore } from './stores';

const homeStore = useHomeStore();
</script>

<style scoped>
@layer page {
  .home-top-bar__actions {
    height: 100%;
    display: flex;
    gap: inherit;
    align-items: inherit;
  }
}
</style>
