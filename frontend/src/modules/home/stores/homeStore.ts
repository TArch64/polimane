import { computed, type Ref, ref } from 'vue';
import { defineStore } from 'pinia';
import type { IFolder, ISchema } from '@/models';
import type {
  IFolderAddSchemaAdapter,
  IFolderUpdateAdapter,
  ISchemaCopyAdapter,
  ISchemaCreateAdapter,
  ISchemaDeleteAdapter,
  ISchemaSelectionAdapter,
  ISchemaUpdateAdapter,
} from './adapters';

export type ListSchema = Omit<ISchema, 'beads' | 'size' | 'screenshotedAt'>;

export interface IListFolder extends IFolder {
  backgroundColor: string | null;
  screenshotPath: string | null;
}

export interface IHomeRouteConfig {
  title?: string;
  selection: ISchemaSelectionAdapter;
  createSchema: ISchemaCreateAdapter;
  updateSchema: ISchemaUpdateAdapter;
  copySchema: ISchemaCopyAdapter;
  deleteSchema: ISchemaDeleteAdapter;
  addSchemaToFolder: IFolderAddSchemaAdapter;
  updateFolder: IFolderUpdateAdapter;
}

export const useHomeStore = defineStore('home', () => {
  const routeConfig = ref<IHomeRouteConfig>({
    selection: {
      ids: new Set(),
      actions: [],
      onClear: null!,
    },

    createSchema: null!,
    updateSchema: null!,
    copySchema: null!,
    deleteSchema: null!,
    addSchemaToFolder: null!,
    updateFolder: null!,
  });

  function setRouteConfig(config: IHomeRouteConfig): void {
    routeConfig.value = config;
  }

  function toConfigRef<N extends keyof IHomeRouteConfig>(name: N): Ref<IHomeRouteConfig[N]> {
    return computed(() => routeConfig.value[name]);
  }

  return {
    setRouteConfig,
    title: toConfigRef('title'),
    selection: toConfigRef('selection'),
    createSchema: toConfigRef('createSchema'),
    updateSchema: toConfigRef('updateSchema'),
    copySchema: toConfigRef('copySchema'),
    deleteSchema: toConfigRef('deleteSchema'),
    addSchemaToFolder: toConfigRef('addSchemaToFolder'),
    updateFolder: toConfigRef('updateFolder'),
  };
});
