<template>
  <div class="schemas-list">
    <HomeSchema
      v-for="schema of schemasStore.schemas"
      :key="schema.id"
      :schema="schema as ISchema"
    />
  </div>

  <div class="schemas-list-loader" v-visible="schemasStore.isLoading">
    <Spinner />
  </div>
</template>

<script setup lang="ts">
import { toRef } from 'vue';
import { useSchemasStore } from '@/modules/home/stores';
import type { ISchema } from '@/models';
import { useInfinityScroll } from '@/composables';
import Spinner from '@/components/Spinner.vue';
import { vVisible } from '@/directives';
import HomeSchema from './HomeSchema.vue';

const schemasStore = useSchemasStore();

useInfinityScroll({
  load: schemasStore.loadNext,
  canLoadNext: toRef(schemasStore, 'canLoadNext'),
});
</script>

<style scoped>
@layer page {
  .schemas-list {
    flex-grow: 1;
    display: grid;
    grid-template-columns: repeat(var(--list-columns), 1fr);
    align-content: start;
    align-items: end;
    gap: 20px;
    padding: 12px;
    --list-columns: 4;
  }

  .schemas-list-loader {
    display: flex;
    justify-content: center;
    padding: 20px 12px;
  }

  @media (max-width: 992px) {
    .schemas-list {
      --list-columns: 3;
    }
  }

  @media (max-width: 768px) {
    .schemas-list {
      --list-columns: 2;
    }
  }
}
</style>
