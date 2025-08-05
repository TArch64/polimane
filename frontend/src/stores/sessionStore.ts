import { defineStore } from 'pinia';
import { computed, nextTick, type Ref, ref } from 'vue';
import type { IUser } from '@/models';
import { useAccessToken, useHttpClient, useRefreshAccessToken } from '@/composables';

export const useSessionStore = defineStore('session', () => {
  const http = useHttpClient();
  const user = ref<IUser | null>(null);
  const accessToken = useAccessToken();
  const refreshAccessToken = useRefreshAccessToken();
  const isLoggedIn = computed(() => !!user.value);

  async function refresh(): Promise<void> {
    try {
      user.value = await http.get<IUser>('/users/current');
    } catch (error) {
      user.value = null;
      console.error(error);
    }
  }

  function setTokens(access: string, refresh: string): void {
    accessToken.value = access;
    refreshAccessToken.value = refresh;
  }

  async function logout(): Promise<void> {
    await http.post('/auth/logout', {});
    accessToken.value = undefined;
    refreshAccessToken.value = undefined;
    await nextTick();
    window.location.reload();
  }

  function updateUser(newUser: Partial<IUser>): void {
    user.value = { ...user.value!, ...newUser };
  }

  return {
    user: user as Ref<IUser>,
    updateUser,
    isLoggedIn,
    refresh,
    setTokens,
    logout,
  };
});
