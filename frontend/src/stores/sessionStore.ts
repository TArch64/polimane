import { defineStore } from 'pinia';
import { computed, type Ref, ref } from 'vue';
import type { IUser } from '@/models';
import { useHttpClient } from '@/composables';

export interface ILoginInput {
  username: string;
  password: string;
}

export const useSessionStore = defineStore('session', () => {
  const httpClient = useHttpClient();
  const user = ref<IUser | null>(null);
  const isLoginedIn = computed(() => !!user.value);

  async function login(input: ILoginInput): Promise<void> {
    user.value = await httpClient.post<IUser, ILoginInput>('/auth/login', input);
  }

  return {
    user: user as Ref<IUser>,
    isLoginedIn,
    login,
  };
});
