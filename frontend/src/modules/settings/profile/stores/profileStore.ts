import { defineStore } from 'pinia';
import { toRef } from 'vue';
import { useSessionStore } from '@/stores';
import type { IUser } from '@/models';
import { type HttpBody, useHttpClient } from '@/composables';

export type ProfileUpdate = Pick<IUser, 'firstName' | 'lastName'>;

export const useProfileStore = defineStore('settings/profile', () => {
  const http = useHttpClient();
  const sessionStore = useSessionStore();

  async function update(data: ProfileUpdate) {
    await http.patch<HttpBody, ProfileUpdate>(['/users/current'], data);
    sessionStore.updateUser(data);
  }

  return {
    user: toRef(sessionStore, 'user'),
    update,
  };
});
