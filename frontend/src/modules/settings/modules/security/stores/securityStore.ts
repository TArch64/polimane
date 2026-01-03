import { defineStore } from 'pinia';
import { useHttpClient } from '@/composables';

export const useSecurityStore = defineStore('settings/security', () => {
  const http = useHttpClient();

  async function resetPassword() {
    await http.post('/users/current/password/reset', {});
  }

  return { resetPassword };
});
