import { defineStore } from 'pinia';
import { computed, type Ref, ref } from 'vue';
import type { IUser } from '@/models';
import { useAuthToken, useHttpClient } from '@/composables';

export interface ILoginInput {
  username: string;
  password: string;
}

interface ILoginResponse {
  user: IUser;
  token: string;
}

export const useSessionStore = defineStore('session', () => {
  const httpClient = useHttpClient();
  const user = ref<IUser | null>(null);
  const authToken = useAuthToken();
  const isLoginedIn = computed(() => !!user.value);

  async function login(input: ILoginInput): Promise<void> {
    const response = await httpClient.post<ILoginResponse, ILoginInput>('/auth/login', input);
    user.value = response.user;
    authToken.value = response.token;
  }

  async function refresh(): Promise<void> {
    try {
      user.value = await httpClient.get<IUser>('/users/current');
    } catch (error) {
      user.value = null;
      console.error(error);
    }
  }

  return {
    user: user as Ref<IUser>,
    isLoginedIn,
    login,
    refresh,
  };
});
