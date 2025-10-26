<template>
  <div class="schemas-list">
    <HomeSchema
      v-for="schema of schemasStore.schemas.data"
      :key="schema.id"
      :schema="schema as ISchema"
    />
  </div>
</template>

<script setup lang="ts">
import { useSchemasStore } from '@/modules/home/stores';
import type { ISchema } from '@/models';
import HomeSchema from './HomeSchema.vue';

const schemasStore = useSchemasStore();
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
