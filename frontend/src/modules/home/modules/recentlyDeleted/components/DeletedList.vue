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
    <HomeListSchema
      disabled
      :ref="itemRef"
      :schema="item"
    />
  </HomeGridList>

  <HomeListLoader :visible="schemasStore.isLoading" />
</template>

<script setup lang="ts">
import {
  HomeGridList,
  HomeListHeading,
  HomeListLoader,
  HomeListSchema,
} from '@/modules/home/components';
import { useDeletedSchemasStore } from '../stores';

const schemasStore = useDeletedSchemasStore();
</script>
