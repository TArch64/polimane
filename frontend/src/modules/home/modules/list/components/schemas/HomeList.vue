<template>
  <div class="schemas-list">
    <CursorSelection
      :list="schemasStore.schemas"
      v-model="schemasStore.selected"
      v-slot="{ item, itemRef }"
    >
      <HomeSchema :ref="itemRef" :schema="item" />
    </CursorSelection>
  </div>

  <div class="schemas-list-loader" v-visible="listStore.list.isLoading">
    <Spinner />
  </div>
</template>

<script setup lang="ts">
import { toRef } from 'vue';
import { useInfinityScroll } from '@/composables';
import { CursorSelection } from '@/components/selection';
import Spinner from '@/components/Spinner.vue';
import { vVisible } from '@/directives';
import { useHomeListStore, useSchemasStore } from '../../stores';
import HomeSchema from './HomeSchema.vue';

const listStore = useHomeListStore();
const schemasStore = useSchemasStore();

useInfinityScroll({
  load: listStore.loadNext,
  canLoadNext: toRef(listStore, 'canLoadNext'),
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
