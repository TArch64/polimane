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
    <DeletedSchema

      :ref="itemRef"
      :schema="item"
    />
  </HomeGridList>

  <HomeListLoader :visible="schemasStore.isLoading" />
</template>

<script setup lang="ts">
import { HomeGridList, HomeListHeading, HomeListLoader } from '@/modules/home/components';
import { useDeletedSchemasStore } from '../stores';
import DeletedSchema from './DeletedSchema.vue';

const schemasStore = useDeletedSchemasStore();
</script>
