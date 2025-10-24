import { defineStore } from 'pinia';
import { computed } from 'vue';
import { type HttpBody, useAsyncData, useHttpClient } from '@/composables';
import type { ISchemaUser, ISchemaUserInvitation } from '@/models';
import type { UrlPath } from '@/helpers';
import { AccessLevel } from '@/enums';
import { useEditorStore } from './editorStore';

interface ISchemaUserList {
  users: ISchemaUser[];
  invitations: ISchemaUserInvitation[];
}

interface IAddUserResponse {
  user?: ISchemaUser;
  invitation?: ISchemaUserInvitation;
}

interface IAddUserBody {
  email: string;
}

interface IUpdateAccessBody {
  access: AccessLevel;
}

export const useSchemaUsersStore = defineStore('schemas/editor/users', () => {
  const editorStore = useEditorStore();
  const baseUrl = computed(() => ['/schemas', editorStore.schema.id, 'users'] as const satisfies UrlPath);
  const http = useHttpClient();

  const list = useAsyncData({
    async loader() {
      const { users, invitations } = await http.get<ISchemaUserList>(baseUrl.value);
      return { users, invitations: invitations ?? [] };
    },

    once: true,

    default: {
      users: [],
      invitations: [],
    },
  });

  const users = computed(() => list.data.users);
  const invitations = computed(() => list.data.invitations);

  async function addUser(email: string): Promise<IAddUserResponse> {
    const response = await http.post<IAddUserResponse, IAddUserBody>(baseUrl.value, { email });
    if (response.user) {
      list.data.users = [...users.value, response.user];
    }
    if (response.invitation) {
      list.data.invitations = [...invitations.value, response.invitation];
    }
    return response;
  }

  async function deleteUser(deletingUser: ISchemaUser): Promise<void> {
    await http.delete([...baseUrl.value, deletingUser.id]);
    list.data.users = users.value.filter((user) => user.id !== deletingUser.id);
  }

  async function updateUserAccess(user: ISchemaUser, access: AccessLevel): Promise<void> {
    await http.patch<HttpBody, IUpdateAccessBody>([...baseUrl.value, user.id, 'access'], {
      access,
    });

    user.access = access;
  }

  return {
    users,
    invitations,
    load: list.load,
    deleteUser,
    addUser,
    updateUserAccess,
  };
});
