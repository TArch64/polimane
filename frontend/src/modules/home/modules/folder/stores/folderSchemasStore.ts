import { computed, type Ref, ref, toRef } from 'vue';
import { defineStore } from 'pinia';
import { type HttpBody, useAsyncData, useHttpClient } from '@/composables';
import type {
  IDeleteManySchemasRequest,
  ISchemaCreateRequest,
  ListSchema,
} from '@/modules/home/stores';
import { AccessLevel } from '@/enums';
import type { SchemaUpdate } from '@/models';

const PAGINATION_PAGE = 100;

interface IListResponse {
  schemas: ListSchema[];
  total: number;
}

type ListRequestParams = {
  folder: string;
  offset: number;
  limit: number;
};

export const useFolderSchemasStore = defineStore('home/folder/schemas', () => {
  const http = useHttpClient();
  const folderId = ref<string>('');

  const list = useAsyncData({
    async loader(current): Promise<IListResponse> {
      return http.get<IListResponse, ListRequestParams>(['/schemas'], {
        folder: folderId.value,
        limit: PAGINATION_PAGE,
        offset: current.schemas.length,
      });
    },

    default: {
      schemas: [],
      total: 0,
    },
  });

  const schemas = computed(() => list.data.schemas);
  const canLoadNext = computed(() => schemas.value.length < list.data.total);

  const selected: Ref<Set<string>> = ref(new Set());
  const clearSelection = () => selected.value = new Set();

  async function load(id: string): Promise<void> {
    folderId.value = id;
    await list.load();
  }

  async function loadNext(): Promise<void> {
    if (list.isLoading) return;
    if (!canLoadNext.value) return;
    return list.load();
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

  async function createSchema(input: ISchemaCreateRequest): Promise<ListSchema> {
    const item = await http.post<ListSchema, ISchemaCreateRequest>('/schemas', input);
    list.data.total++;
    return item;
  }

  async function deleteMany(ids: string[]): Promise<void> {
    const idsSet = new Set(ids);

    await list.optimisticUpdate()
      .inTransition()
      .begin((state) => {
        state.schemas = state.schemas.filter((schema) => !idsSet.has(schema.id));
        state.total -= ids.length;
      })
      .commit(async () => {
        await http.delete<HttpBody, IDeleteManySchemasRequest>(['/schemas', 'delete-many'], { ids });
      });

    if (canLoadNext.value && schemas.value.length < PAGINATION_PAGE) {
      await loadNext();
    }
  }

  async function deleteSchema(deleting: ListSchema): Promise<void> {
    return deleteMany([deleting.id]);
  }

  async function copySchema(copyingSchema: ListSchema): Promise<ListSchema> {
    const item = await http.post<ListSchema>(['/schemas', copyingSchema.id, 'copy'], {});
    list.data.total++;
    return item;
  }

  async function updateSchema(updatingSchema: ListSchema, patch: SchemaUpdate): Promise<void> {
    await http.patch<HttpBody, SchemaUpdate>(['/schemas', updatingSchema.id], patch);
    Object.assign(updatingSchema, patch);
  }

  return {
    schemas,
    selected,
    canLoadNext,
    isLoading: toRef(list, 'isLoading'),
    load,
    loadNext,
    clearSelection,
    filterIdsByAccess,
    createSchema,
    deleteMany,
    deleteSchema,
    copySchema,
    updateSchema,
  };
});
