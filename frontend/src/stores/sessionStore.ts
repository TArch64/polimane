import { defineStore } from 'pinia';
import { nextTick, type Ref, ref } from 'vue';
import type { IUser } from '@/models';
import { useAuthorized, useHttpClient } from '@/composables';

export const useSessionStore = defineStore('session', () => {
  const http = useHttpClient();
  const user = ref<IUser | null>(null);
  const authorized = useAuthorized();

  async function refresh(): Promise<void> {
    try {
      user.value = await http.get<IUser>('/users/current');
      authorized.value = true;
    } catch (error) {
      user.value = null;
      authorized.value = false;
      console.error(error);
    }
  }

  async function onLogout(): Promise<void> {
    authorized.value = false;
    await nextTick();
    window.location.reload();
  }

  async function logout(): Promise<void> {
    await http.post('/auth/logout', {});
    await onLogout();
  }

  function updateUser(newUser: Partial<IUser>): void {
    user.value = { ...user.value!, ...newUser };
  }

  return {
    user: user as Ref<IUser>,
    updateUser,
    isLoggedIn: authorized,
    refresh,
    logout,
    onLogout,
  };
});
