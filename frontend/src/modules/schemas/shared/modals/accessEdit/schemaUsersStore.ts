import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import { type HttpBody, useAsyncData, useHttpClient } from '@/composables';
import type { ISchemaUser, ISchemaUserInvitation } from '@/models';
import type { UrlPath } from '@/helpers';
import { AccessLevel } from '@/enums';

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

type IDeleteInvitationParams = {
  email: string;
};

interface IUpdateInvitationAccessBody extends IUpdateAccessBody {
  email: string;
}

export const useSchemaUsersStore = defineStore('schemas/users', () => {
  const schemaId = ref<string>('');
  const baseUrl = computed(() => ['/schemas', schemaId.value, 'users'] as const satisfies UrlPath);
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

  async function load(_schemaId: string): Promise<void> {
    if (_schemaId !== schemaId.value) {
      schemaId.value = _schemaId;
      list.reset();
    }

    await list.load();
  }

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

  async function deleteInvitation(deletingInvitation: ISchemaUserInvitation): Promise<void> {
    await http.delete<HttpBody, IDeleteInvitationParams>([...baseUrl.value, 'invitations'], {
      email: deletingInvitation.email,
    });

    list.data.invitations = invitations.value.filter((invitation) => {
      return invitation.email !== deletingInvitation.email;
    });
  }

  async function updateInvitationAccess(invitation: ISchemaUserInvitation, access: AccessLevel): Promise<void> {
    await http.patch<HttpBody, IUpdateInvitationAccessBody>([...baseUrl.value, 'invitations', 'access'], {
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
