import { computed, ref } from 'vue';
import { defineStore } from 'pinia';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import type { IFolder, ISchema } from '@/models';
import type { IFolderAddSchemaStrategy, ISchemaCreateStrategy } from './strategies';

export type ListSchema = Omit<ISchema, 'beads' | 'size' | 'screenshotedAt'>;

export interface IListFolder extends IFolder {
  backgroundColor: string | null;
  screenshotPath: string | null;
}

export interface IHomeSelectionState {
  count: number;
  title: string;
  actions: MaybeContextMenuAction[];
  onClear: () => void;
}

export interface IHomeStrategies {
  createSchema: ISchemaCreateStrategy;
  addSchemaToFolder: IFolderAddSchemaStrategy;
}

export const useHomeStore = defineStore('home', () => {
  const selection = ref<IHomeSelectionState | null>(null);
  const setSelection = (state: IHomeSelectionState | null) => selection.value = state;

  const strategies = ref<IHomeStrategies | null>(null);
  const setStrategies = (value: IHomeStrategies | null) => strategies.value = value;
  const createSchema = computed(() => strategies.value?.createSchema || null);
  const addSchemaToFolder = computed(() => strategies.value?.addSchemaToFolder || null);

  return {
    selection,
    createSchema,
    addSchemaToFolder,
    setSelection,
    setStrategies,
  };
});
