import { defineStore } from 'pinia';
import { computed, toRef } from 'vue';
import { useAsyncData, useHttpClient } from '@/composables';
import type { ISchemaUser } from '@/models';
import type { UrlPath } from '@/helpers';
import { useEditorStore } from './editorStore';

interface IAddUserResponse {
  invited: boolean;
  user?: ISchemaUser;
}

interface IAddUserBody {
  email: string;
}

export const useSchemaUsersStore = defineStore('schemas/editor/users', () => {
  const editorStore = useEditorStore();
  const baseUrl = computed(() => ['/schemas', editorStore.schema.id, 'users'] as const satisfies UrlPath);
  const http = useHttpClient();

  const users = useAsyncData({
    loader: async () => http.get<ISchemaUser[]>(baseUrl.value),
    once: true,
    default: [],
  });

  async function addUser(email: string): Promise<IAddUserResponse> {
    const response = await http.post<IAddUserResponse, IAddUserBody>(baseUrl.value, { email });
    if (response.user) {
      users.data = [...users.data, response.user];
    }
    return response;
  }

  async function deleteUser(deletingUser: ISchemaUser): Promise<void> {
    await http.delete([...baseUrl.value, deletingUser.id]);
    users.data = users.data.filter((user) => user.id !== deletingUser.id);
  }

  return {
    users: toRef(users, 'data'),
    load: users.load,
    deleteUser,
    addUser,
  };
});
