import { defineStore } from 'pinia';
import { computed, ref, type Ref, toRef } from 'vue';
import {
  type HttpBody,
  type IOptimisticOptions,
  useAsyncData,
  useHttpClient,
} from '@/composables';
import type { IFolder, ISchema, SchemaUpdate } from '@/models';
import { AccessLevel } from '@/enums';

const PAGINATION_PAGE = 100;

export type ListSchema = Omit<ISchema, 'beads' | 'size'>;

interface IListResponse {
  folders: IFolder[];
  schemas: ListSchema[];
  total: number;
}

type ListRequestParams = {
  offset: number;
  limit: number;
};

export interface ICreateSchemaRequest {
  name: string;
}

interface IDeleteManySchemasBody {
  ids: string[];
}

export const useSchemasStore = defineStore('schemas/list', () => {
  const http = useHttpClient();

  const list = useAsyncData({
    loader: async (current): Promise<IListResponse> => {
      const response = await http.get<IListResponse, ListRequestParams>('/schemas', {
        limit: PAGINATION_PAGE,
        offset: current.folders.length + current.schemas.length,
      });

      return {
        folders: [...current.folders, ...response.folders],
        schemas: [...current.schemas, ...response.schemas],
        total: response.total,
      };
    },

    default: {
      folders: [],
      schemas: [],
      total: 0,
    },
  });

  const schemas = computed(() => list.data.schemas);
  const hasSchemas = computed(() => !!schemas.value.length);
  const canLoadNext = computed(() => schemas.value.length < list.data.total);

  const selected: Ref<Set<string>> = ref(new Set());
  const clearSelection = () => selected.value = new Set();

  function load(): Promise<void> {
    list.reset();
    return list.load();
  }

  async function loadNext(): Promise<void> {
    if (list.isLoading) return;
    if (!canLoadNext.value) return;
    return list.load();
  }

  async function createSchema(input: ICreateSchemaRequest): Promise<ListSchema> {
    const item = await http.post<ListSchema, ICreateSchemaRequest>('/schemas', input);
    list.data.total++;
    return item;
  }

  function filterIdsByAccess(ids: Set<string>, access: AccessLevel): string[] {
    const result = new Set<string>();

    for (const schema of schemas.value) {
      if (ids.has(schema.id) && schema.access >= access) {
        result.add(schema.id);
      }
    }

    return [...result];
  }

  async function deleteMany(ids: string[]): Promise<void> {
    const optimisticOptions: IOptimisticOptions = { transition: true };
    const idsSet = new Set(ids);

    list.makeOptimisticUpdate(({ schemas, total, ...rest }) => ({
      ...rest,
      schemas: schemas.filter((schema) => !idsSet.has(schema.id)),
      total: total - ids.length,
    }), optimisticOptions);

    await list.executeOptimisticUpdate(async () => {
      await http.delete<HttpBody, IDeleteManySchemasBody>(['/schemas', 'delete-many'], { ids });
    }, optimisticOptions);

    if (canLoadNext.value && schemas.value.length < PAGINATION_PAGE) {
      await loadNext();
    }
  }

  async function deleteSchema(deleting: ListSchema): Promise<void> {
    return deleteMany([deleting.id]);
  }

  async function copySchema(copyingSchema: ListSchema): Promise<ListSchema> {
    const item = await http.post<ListSchema, HttpBody>(['/schemas', copyingSchema.id, 'copy'], {});
    list.data.total++;
    return item;
  }

  async function updateSchema(updatingSchema: ListSchema, patch: SchemaUpdate): Promise<void> {
    await http.patch<HttpBody, SchemaUpdate>(['/schemas', updatingSchema.id], patch);
    Object.assign(updatingSchema, patch);
  }

  return {
    schemas,
    hasSchemas,
    isLoading: toRef(list, 'isLoading'),
    canLoadNext,
    load,
    loadNext,
    selected,
    clearSelection,
    createSchema,
    deleteSchema,
    deleteMany,
    copySchema,
    updateSchema,
    filterIdsByAccess,
  };
});
