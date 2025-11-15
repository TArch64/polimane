import { defineStore } from 'pinia';
import { computed, ref, type Ref } from 'vue';
import { type HttpBody, type IOptimisticOptions, useHttpClient } from '@/composables';
import type { SchemaUpdate } from '@/models';
import { AccessLevel } from '@/enums';
import type { ListSchema } from '@/modules/home/stores';
import { useHomeListStore } from './homeListStore';

const PAGINATION_PAGE = 100;

export interface ICreateSchemaRequest {
  name: string;
}

interface IDeleteManySchemasBody {
  ids: string[];
}

export const useSchemasStore = defineStore('schemas/list/schemas', () => {
  const http = useHttpClient();
  const listStore = useHomeListStore();

  const schemas = computed(() => listStore.list.data.schemas);
  const hasSchemas = computed(() => !!schemas.value.length);

  const selected: Ref<Set<string>> = ref(new Set());
  const clearSelection = () => selected.value = new Set();

  async function createSchema(input: ICreateSchemaRequest): Promise<ListSchema> {
    const item = await http.post<ListSchema, ICreateSchemaRequest>('/schemas', input);
    listStore.list.data.total++;
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

    listStore.list.makeOptimisticUpdate(({ schemas, total, ...rest }) => ({
      ...rest,
      schemas: schemas.filter((schema) => !idsSet.has(schema.id)),
      total: total - ids.length,
    }), optimisticOptions);

    await listStore.list.executeOptimisticUpdate(async () => {
      await http.delete<HttpBody, IDeleteManySchemasBody>(['/schemas', 'delete-many'], { ids });
    }, optimisticOptions);

    if (listStore.canLoadNext && schemas.value.length < PAGINATION_PAGE) {
      await listStore.loadNext();
    }
  }

  async function deleteSchema(deleting: ListSchema): Promise<void> {
    return deleteMany([deleting.id]);
  }

  async function copySchema(copyingSchema: ListSchema): Promise<ListSchema> {
    const item = await http.post<ListSchema, HttpBody>(['/schemas', copyingSchema.id, 'copy'], {});
    listStore.list.data.total++;
    return item;
  }

  async function updateSchema(updatingSchema: ListSchema, patch: SchemaUpdate): Promise<void> {
    await http.patch<HttpBody, SchemaUpdate>(['/schemas', updatingSchema.id], patch);
    Object.assign(updatingSchema, patch);
  }

  return {
    schemas,
    hasSchemas,
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
