import { computed, ref } from 'vue';
import { defineStore } from 'pinia';
import type { MaybeContextMenuAction } from '@/components/contextMenu';
import type { IFolder, ISchema } from '@/models';
import type {
  IFolderAddSchemaStrategy,
  ISchemaCopyStrategy,
  ISchemaCreateStrategy,
  ISchemaDeleteStrategy,
  ISchemaUpdateStrategy,
} from './strategies';

export type ListSchema = Omit<ISchema, 'beads' | 'size' | 'screenshotedAt'>;

export interface IListFolder extends IFolder {
  backgroundColor: string | null;
  screenshotPath: string | null;
}

export interface IHomeSelectionState {
  title: string;
  ids: Set<string>;
  actions: MaybeContextMenuAction[];
  onClear: () => void;
}

export interface IHomeRouteConfig {
  title: string;
  selection: IHomeSelectionState;
  createSchema: ISchemaCreateStrategy;
  updateSchema: ISchemaUpdateStrategy;
  copySchema: ISchemaCopyStrategy;
  deleteSchema: ISchemaDeleteStrategy;
  addSchemaToFolder: IFolderAddSchemaStrategy;
}

export const useHomeStore = defineStore('home', () => {
  const routeConfig = ref<Partial<IHomeRouteConfig> | undefined>(undefined);

  function setRouteConfig(config: Partial<IHomeRouteConfig> | undefined): void {
    routeConfig.value = config;
  }

  const title = computed(() => routeConfig.value?.title || '');
  const selection = computed(() => routeConfig.value?.selection);
  const createSchema = computed(() => routeConfig.value?.createSchema);
  const updateSchema = computed(() => routeConfig.value?.updateSchema);
  const copySchema = computed(() => routeConfig.value?.copySchema);
  const deleteSchema = computed(() => routeConfig.value?.deleteSchema);
  const addSchemaToFolder = computed(() => routeConfig.value?.addSchemaToFolder);

  return {
    setRouteConfig,
    title,
    selection,
    createSchema,
    updateSchema,
    copySchema,
    deleteSchema,
    addSchemaToFolder,
  };
});
