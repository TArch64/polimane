import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import { type HttpBody, useAsyncData, useHttpClient } from '@/composables';
import type { ISchemaUser, ISchemaUserInvitation } from '@/models';
import type { UrlPath } from '@/helpers';
import { AccessLevel } from '@/enums';

type SchemaIdParams = {
  ids: string[];
};

interface ISchemaUserList {
  users: ISchemaUser[];
  invitations: ISchemaUserInvitation[];
}

interface IAddUserResponse {
  user?: ISchemaUser;
  invitation?: ISchemaUserInvitation;
}

interface IAddUserBody extends SchemaIdParams {
  email: string;
}

interface IUpdateAccessBody extends SchemaIdParams {
  access: AccessLevel;
}

interface IDeleteInvitationBody extends SchemaIdParams {
  email: string;
}

interface IUpdateInvitationAccessBody extends IUpdateAccessBody {
  email: string;
}

export const useSchemaUsersStore = defineStore('schemas/users', () => {
  const schemaIds = ref<string[]>([]);
  const baseUrl = computed(() => ['/schemas', 'users'] as const satisfies UrlPath);
  const http = useHttpClient();

  const list = useAsyncData({
    async loader() {
      const {
        users,
        invitations,
      } = await http.get<ISchemaUserList, SchemaIdParams>(baseUrl.value, {
        ids: schemaIds.value,
      });

      return {
        users,
        invitations: invitations ?? [],
      };
    },

    default: {
      users: [],
      invitations: [],
    },
  });

  async function load(ids: string[]): Promise<void> {
    const existing = new Set(schemaIds.value);
    const loading = new Set(ids);

    if (existing.symmetricDifference(loading).size) {
      schemaIds.value = Array.from(loading);
      await list.load();
    }
  }

  const users = computed(() => list.data.users);
  const invitations = computed(() => list.data.invitations);

  async function addUser(email: string): Promise<IAddUserResponse> {
    const response = await http.post<IAddUserResponse, IAddUserBody>(baseUrl.value, {
      ids: schemaIds.value,
      email,
    });

    if (response.user && !users.value.find((user) => user.email === email)) {
      list.data.users = [...users.value, response.user];
    }
    if (response.invitation && !invitations.value.find((invitation) => invitation.email === email)) {
      list.data.invitations = [...invitations.value, response.invitation];
    }
    return response;
  }

  async function deleteUser(deletingUser: ISchemaUser): Promise<void> {
    list.makeOptimisticUpdate((current) => ({
      ...current,
      users: current.users.filter((user) => user.id !== deletingUser.id),
    }));

    await list.executeOptimisticUpdate(async () => {
      await http.delete<HttpBody, SchemaIdParams>([...baseUrl.value, deletingUser.id], {
        ids: schemaIds.value,
      });
    });
  }

  async function updateUserAccess(user: ISchemaUser, access: AccessLevel): Promise<void> {
    await http.patch<HttpBody, IUpdateAccessBody>([...baseUrl.value, user.id, 'access'], {
      ids: schemaIds.value,
      access,
    });

    user.access = access;
  }

  async function deleteInvitation(deletingInvitation: ISchemaUserInvitation): Promise<void> {
    await http.delete<HttpBody, IDeleteInvitationBody>([...baseUrl.value, 'invitations'], {
      ids: schemaIds.value,
      email: deletingInvitation.email,
    });

    list.data.invitations = invitations.value.filter((invitation) => {
      return invitation.email !== deletingInvitation.email;
    });
  }

  async function updateInvitationAccess(invitation: ISchemaUserInvitation, access: AccessLevel): Promise<void> {
    await http.patch<HttpBody, IUpdateInvitationAccessBody>([...baseUrl.value, 'invitations', 'access'], {
      ids: schemaIds.value,
      email: invitation.email,
      access,
    });

    invitation.access = access;
  }

  return {
    users,
    invitations,
    load,
    addUser,
    deleteUser,
    updateUserAccess,
    deleteInvitation,
    updateInvitationAccess,
  };
});
