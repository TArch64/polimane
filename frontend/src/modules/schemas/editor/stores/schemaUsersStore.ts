import { defineStore } from 'pinia';
import { toRef } from 'vue';
import { useAsyncData, useHttpClient } from '@/composables';
import type { ISchemaUser } from '@/models';
import type { UrlPath } from '@/helpers';
import { useEditorStore } from './editorStore';

export const useSchemaUsersStore = defineStore('schemas/editor/users', () => {
  const editorStore = useEditorStore();
  const http = useHttpClient();

  const users = useAsyncData({
    async loader(): Promise<ISchemaUser[]> {
      const url: UrlPath = ['/schemas', editorStore.schema.id, 'users'];
      return http.get<ISchemaUser[]>(url);
    },

    default: [],
  });

  async function load() {
    if (users.isInitial) await users.load();
  }

  return {
    load,
    users: toRef(users, 'data'),
  };
});
