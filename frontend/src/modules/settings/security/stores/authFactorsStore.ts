import { ref } from 'vue';
import { defineStore } from 'pinia';
import { useHttpClient } from '@/composables';
import type { IAuthFactor } from '../models';

export const useAuthFactorsStore = defineStore('settings/auth-factors', () => {
  const http = useHttpClient();
  const list = ref<IAuthFactor[]>([]);

  async function load(): Promise<void> {
    list.value = await http.get<IAuthFactor[]>('/users/current/auth-factors');
  }

  return { list, load };
});
