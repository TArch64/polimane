import { ref } from 'vue';
import { defineStore } from 'pinia';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import type { ISchema } from '@/models';
import type { ISchemaCreateStrategy } from './strategies';

export type ListSchema = Omit<ISchema, 'beads' | 'size'>;

export interface IHomeSelectionState {
  count: number;
  title: string;
  actions: MaybeContextMenuAction[];
  onClear: () => void;
}

export const useHomeStore = defineStore('home', () => {
  const selection = ref<IHomeSelectionState | null>(null);
  const setSelection = (state: IHomeSelectionState | null) => selection.value = state;

  const createSchema = ref<ISchemaCreateStrategy | null>(null);
  const setCreateSchemaStrategy = (strategy: ISchemaCreateStrategy | null) => createSchema.value = strategy;

  return {
    selection,
    createSchema,
    setSelection,
    setCreateSchemaStrategy,
  };
});
