<template>
  <HomeListCard
    :to="folderRoute"
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
import { computed } from 'vue';
import type { RouteLocationRaw } from 'vue-router';
import type { IListFolder } from '@/modules/home/stores';
import { DEFAULT_SCHEMA_BACKGROUND } from '@/config';
import { HomeListCard, HomeListScreenshot } from '@/modules/home/components';

const props = defineProps<{
  folder: IListFolder;
}>();

const folderRoute = computed((): RouteLocationRaw => ({
  name: 'home-folder',
  params: { folderId: props.folder.id },
}));
</script>
