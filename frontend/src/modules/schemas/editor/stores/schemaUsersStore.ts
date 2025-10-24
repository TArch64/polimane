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
    once: true,
    default: [],
  });

  async function deleteUser(deletingUser: ISchemaUser): Promise<void> {
    await http.delete(['/schemas', editorStore.schema.id, 'users', deletingUser.id]);
    users.data = users.data.filter((user) => user.id !== deletingUser.id);
  }

  return {
    load: users.load,
    deleteUser,
    users: toRef(users, 'data'),
  };
});
