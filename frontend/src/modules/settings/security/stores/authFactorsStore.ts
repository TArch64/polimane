import { ref } from 'vue';
import { defineStore } from 'pinia';
import { useHttpClient } from '@/composables';
import type { IAuthFactor, IAuthFactorInit } from '../models';

interface ICreateAuthFactorBody {
  challengeId: string;
  code: string;
}

export const useAuthFactorsStore = defineStore('settings/auth-factors', () => {
  const http = useHttpClient();
  const list = ref<IAuthFactor[]>([]);

  async function load(): Promise<void> {
    list.value = await http.get<IAuthFactor[]>('/users/current/auth-factors');
  }

  async function initNewFactor(): Promise<IAuthFactorInit> {
    return http.post('/users/current/auth-factors/init', {});
  }

  function createFactor(challengeId: string, code: string): Promise<IAuthFactor> {
    return http.post<IAuthFactor, ICreateAuthFactorBody>('/users/current/auth-factors', {
      challengeId,
      code,
    });
  }

  function addFactor(factor: IAuthFactor): void {
    list.value.push(factor);
  }

  async function deleteFactor(deleting: IAuthFactor): Promise<void> {
    await http.delete(['/users/current/auth-factors', deleting.id]);
    list.value = list.value.filter((factor) => factor.id !== deleting.id);
  }

  return { list, load, initNewFactor, createFactor, addFactor, deleteFactor };
});
