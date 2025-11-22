<template>
  <template v-if="foldersStore.hasFolders">
    <HomeListHeading>
      Директорії
    </HomeListHeading>

    <HomeGridList
      :list="foldersStore.folders"
      class="home-list__folders"
      v-slot="{ item }"
    >
      <HomeFolder :folder="item" />
    </HomeGridList>

    <HomeListHeading>
      Схеми для Бісеру
    </HomeListHeading>
  </template>

  <HomeGridList
    selectable
    :list="schemasStore.schemas"
    v-model:selected="schemasStore.selected"
    v-slot="{ item, itemRef }"
  >
    <HomeListSchema :ref="itemRef" :schema="item" />
  </HomeGridList>

  <HomeListLoader :visible="listStore.list.isLoading" />
</template>

<script setup lang="ts">
import { toRef } from 'vue';
import { useInfinityScroll } from '@/composables';
import {
  HomeGridList,
  HomeListHeading,
  HomeListLoader,
  HomeListSchema,
} from '@/modules/home/components';
import { useFoldersStore, useHomeListStore, useSchemasStore } from '../stores';
import HomeFolder from './HomeFolder.vue';

const listStore = useHomeListStore();
const schemasStore = useSchemasStore();
const foldersStore = useFoldersStore();

useInfinityScroll({
  load: listStore.loadNext,
  canLoadNext: toRef(listStore, 'canLoadNext'),
});
</script>

<style scoped>
@layer page {
  .home-list__folders {
    margin-bottom: 24px;
  }
}
</style>
