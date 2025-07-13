import { defineStore } from 'pinia';
import { ref } from 'vue';
import { useHttpClient } from '@/composables';

type ILoginUrlParams = {
  'return-to': string;
};

interface ILoginUrlResponse {
  url: string;
}

export const useAuthStore = defineStore('auth', () => {
  const http = useHttpClient();
  const loginUrl = ref('');

  async function load(): Promise<void> {
    const { url } = await http.get<ILoginUrlResponse, ILoginUrlParams>('/auth/login');

    loginUrl.value = url;
  }

  return { load, loginUrl };
});
