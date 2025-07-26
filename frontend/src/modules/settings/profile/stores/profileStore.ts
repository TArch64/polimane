import { defineStore } from 'pinia';
import { ref, toRef } from 'vue';
import { useSessionStore } from '@/stores';
import type { IUser } from '@/models';
import { type HttpBody, useHttpClient } from '@/composables';

export type ProfileUpdate = Partial<Pick<IUser, 'firstName' | 'lastName' | 'email'>>;

interface IVerifyEmailBody {
  code: string;
}

export const useProfileStore = defineStore('settings/profile', () => {
  const http = useHttpClient();
  const sessionStore = useSessionStore();

  const isChangeVerifyingEmail = ref(false);
  const activateChangeVerifyingEmail = () => isChangeVerifyingEmail.value = true;

  async function update(data: ProfileUpdate) {
    await http.patch<HttpBody, ProfileUpdate>(['/users/current'], data);

    sessionStore.updateUser({
      ...data,
      isEmailVerified: !data.email,
    });

    isChangeVerifyingEmail.value = false;
  }

  async function retryEmailVerification() {
    await http.post(['/users/current/email/verify/retry'], {});
  }

  async function verifyEmail(code: string) {
    await http.post<HttpBody, IVerifyEmailBody>(['/users/current/email/verify'], {
      code,
    });

    sessionStore.updateUser({ isEmailVerified: true });
    isChangeVerifyingEmail.value = false;
  }

  return {
    user: toRef(sessionStore, 'user'),
    isChangeVerifyingEmail,
    activateChangeVerifyingEmail,
    update,
    retryEmailVerification,
    verifyEmail,
  };
});
