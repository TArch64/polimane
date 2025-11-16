<template>
  <HomeListCard
    to="/"
    :menu-title="folder.name"
    :menu-actions="[]"
  >
    {{ folder.name }}
  </HomeListCard>
</template>

<script setup lang="ts">
import type { IFolder } from '@/models';
import HomeListCard from './HomeListCard.vue';

defineProps<{
  folder: IFolder;
}>();
</script>

<style scoped>
@layer page {

}
</style>
