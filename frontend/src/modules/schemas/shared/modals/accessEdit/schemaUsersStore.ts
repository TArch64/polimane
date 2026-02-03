import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import { type HttpBody, useAsyncData, useHttpClient } from '@/composables';
import type { ISchemaUser, ISchemaUserInvitation } from '@/models';
import { type UrlPath } from '@/helpers';
import { AccessLevel } from '@/enums';
import { useSchemasSharedAccessCounter } from '@/composables/subscription';

type SchemaIdsParams = {
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

interface IAddUserBody extends SchemaIdsParams {
  email: string;
}

interface IUpdateAccessBody extends SchemaIdsParams {
  access: AccessLevel;
}

interface IDeleteInvitationBody extends SchemaIdsParams {
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
      } = await http.get<ISchemaUserList, SchemaIdsParams>(baseUrl.value, {
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

  const counter = useSchemasSharedAccessCounter(() => ({
    counters: {
      schemaBeads: 0,
      sharedAccess: users.value.length + invitations.value.length,
    },
  }));

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

  async function deleteUser(deleting: ISchemaUser): Promise<void> {
    await list.optimisticUpdate()
      .begin((state) => {
        state.users = state.users.filter((user) => user.id !== deleting.id);
      })
      .commit(async () => {
        await http.delete<HttpBody, SchemaIdsParams>([...baseUrl.value, deleting.id], {
          ids: schemaIds.value,
        });
      });
  }

  async function updateUserAccess(updating: ISchemaUser, access: AccessLevel): Promise<void> {
    await list.optimisticUpdate()
      .begin((state) => {
        const user = state.users.find((user) => user.id === updating.id)!;
        user.access = access;
        user.isEvenAccess = true;
        user.isAllAccess = true;
      })
      .commit(async () => {
        await http.patch<HttpBody, IUpdateAccessBody>([...baseUrl.value, updating.id, 'access'], {
          ids: schemaIds.value,
          access,
        });
      });
  }

  async function deleteInvitation(deleting: ISchemaUserInvitation): Promise<void> {
    await list.optimisticUpdate()
      .begin((state) => {
        state.invitations = state.invitations.filter((invitation) => invitation.email !== deleting.email);
      })
      .commit(async () => {
        await http.delete<HttpBody, IDeleteInvitationBody>([...baseUrl.value, 'invitations'], {
          ids: schemaIds.value,
          email: deleting.email,
        });
      });
  }

  async function updateInvitationAccess(updating: ISchemaUserInvitation, access: AccessLevel): Promise<void> {
    await list.optimisticUpdate()
      .begin((state) => {
        const invitation = state.invitations.find((invitation) => invitation.email === updating.email)!;
        invitation.access = access;
        invitation.isEvenAccess = true;
        invitation.isAllAccess = true;
      })
      .commit(async () => {
        await http.patch<HttpBody, IUpdateInvitationAccessBody>([...baseUrl.value, 'invitations', 'access'], {
          ids: schemaIds.value,
          email: updating.email,
          access,
        });
      });
  }

  return {
    users,
    invitations,
    counter,
    load,
    addUser,
    deleteUser,
    updateUserAccess,
    deleteInvitation,
    updateInvitationAccess,
  };
});
