import { defineStore } from 'pinia';
import { computed, type Ref, ref } from 'vue';
import type { IUser } from '@/models';
import { useAccessToken, useHttpClient, useRefreshAccessToken } from '@/composables';

interface ILogoutResponse {
  url: string;
}

export const useSessionStore = defineStore('session', () => {
  const httpClient = useHttpClient();
  const user = ref<IUser | null>(null);
  const accessToken = useAccessToken();
  const refreshAccessToken = useRefreshAccessToken();
  const isLoggedIn = computed(() => !!user.value);

  async function refresh(): Promise<void> {
    try {
      user.value = await httpClient.get<IUser>('/users/current');
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
    const { url } = await httpClient.get<ILogoutResponse>(['/auth/logout']);
    accessToken.value = undefined;
    refreshAccessToken.value = undefined;
    window.location.assign(url);
  }

  return {
    user: user as Ref<IUser>,
    isLoggedIn,
    refresh,
    setTokens,
    logout,
  };
});
