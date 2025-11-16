<template>
  <template v-if="foldersStore.hasFolders">
    <HomeListHeading>
      Директорії
    </HomeListHeading>

    <div class="home-list__items home-list__folders">
      <HomeFolder
        v-for="folder of foldersStore.folders"
        :key="folder.id"
        :folder
      />
    </div>

    <HomeListHeading>
      Схеми для Бісеру
    </HomeListHeading>
  </template>

  <div class="home-list__items">
    <CursorSelection
      :list="schemasStore.schemas"
      v-model="schemasStore.selected"
      v-slot="{ item, itemRef }"
    >
      <HomeSchema :ref="itemRef" :schema="item" />
    </CursorSelection>
  </div>

  <div class="home-list__loader" v-visible="listStore.list.isLoading">
    <Spinner />
  </div>
</template>

<script setup lang="ts">
import { toRef } from 'vue';
import { useInfinityScroll } from '@/composables';
import { CursorSelection } from '@/components/selection';
import Spinner from '@/components/Spinner.vue';
import { vVisible } from '@/directives';
import { useFoldersStore, useHomeListStore, useSchemasStore } from '../../stores';
import HomeSchema from './HomeSchema.vue';
import HomeFolder from './HomeFolder.vue';
import HomeListHeading from './HomeListHeading.vue';

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
  .home-list__items {
    flex-grow: 1;
    display: grid;
    grid-template-columns: repeat(var(--list-columns), 1fr);
    align-content: start;
    align-items: end;
    gap: 20px;
    padding: 12px;
    --list-columns: 4;
  }

  .home-list__folders {
    margin-bottom: 24px;
  }

  .home-list__loader {
    display: flex;
    justify-content: center;
    padding: 20px 12px;
  }

  @media (max-width: 992px) {
    .home-list__items {
      --list-columns: 3;
    }
  }

  @media (max-width: 768px) {
    .home-list__items {
      --list-columns: 2;
    }
  }
}
</style>
