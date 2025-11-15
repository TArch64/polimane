import { ref } from 'vue';
import { defineStore } from 'pinia';
import type { MaybeContextMenuAction } from '@/components/contextMenu';

export interface IHomeSelectionState {
  count: number;
  title: string;
  actions: MaybeContextMenuAction[];
  onClear: () => void;
}

export const useHomeStore = defineStore('home', () => {
  const selection = ref<IHomeSelectionState | null>(null);
  const setSelection = (state: IHomeSelectionState | null) => selection.value = state;

  return {
    selection,
    setSelection,
  };
});
