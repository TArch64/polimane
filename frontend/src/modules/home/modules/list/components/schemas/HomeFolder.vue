<template>
  <HomeListCard
    to="/"
    :menu-title="folder.name"
    :menu-actions="[]"
  >
    <HomeListScreenshot
      :path="folder.screenshotPath"
      :alt="`Скріншот схеми ${folder.name}`"
      :background-color="folder.backgroundColor || DEFAULT_SCHEMA_BACKGROUND"
    />

    {{ folder.name }}
  </HomeListCard>
</template>

<script setup lang="ts">
import type { IListFolder } from '@/modules/home/stores';
import { DEFAULT_SCHEMA_BACKGROUND } from '@/config';
import HomeListCard from './HomeListCard.vue';
import HomeListScreenshot from './HomeListScreenshot.vue';

defineProps<{
  folder: IListFolder;
}>();
</script>
