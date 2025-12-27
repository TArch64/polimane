<template>
  <HomeListHeading>
    Схеми для Бісеру
  </HomeListHeading>

  <HomeGridList
    selectable
    :list="schemasStore.schemas"
    v-model:selected="schemasStore.selected"
    v-slot="{ item, itemRef }"
  >
    <FolderSchema :ref="itemRef" :schema="item" />
  </HomeGridList>

  <HomeListLoader :visible="schemasStore.isLoading" />
</template>

<script setup lang="ts">
import { toRef } from 'vue';
import { useInfinityScroll } from '@/composables';
import { HomeGridList, HomeListHeading, HomeListLoader } from '@/modules/home/components';
import { useFolderSchemasStore } from '../stores';
import FolderSchema from './FolderSchema.vue';

const schemasStore = useFolderSchemasStore();

useInfinityScroll({
  load: schemasStore.loadNext,
  canLoadNext: toRef(schemasStore, 'canLoadNext'),
});
</script>
