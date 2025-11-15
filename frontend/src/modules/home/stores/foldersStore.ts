import { defineStore } from 'pinia';
import { computed } from 'vue';
import { useHomeListStore } from './homeListStore';

export const useFoldersStore = defineStore('home/list/folders', () => {
  const listStore = useHomeListStore();

  const folders = computed(() => listStore.list.data.folders);
  const hasFolders = computed(() => !!folders.value.length);

  return {
    folders,
    hasFolders,
  };
});
