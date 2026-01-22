import { defineStore } from 'pinia';
import { nextTick, type Ref, ref, watch } from 'vue';
import type { IUser } from '@/models';
import { UpdateCountersMiddleware, useAuthorized, useHttpClient } from '@/composables';

export const useSessionStore = defineStore('session', () => {
  const http = useHttpClient();
  const updateCountersMiddleware = http.getMiddleware(UpdateCountersMiddleware)!;

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
    if (user.value) {
      user.value = { ...user.value, ...newUser };
    }
  }

  watch(
    user,
    (user) => updateCountersMiddleware.user = user,
    { immediate: true },
  );

  updateCountersMiddleware.onUserUpdate.listen((counters) => {
    updateUser({
      subscription: {
        ...user.value!.subscription,
        counters: counters,
      },
    });
  });

  return {
    user: user as Ref<IUser>,
    updateUser,
    isLoggedIn: authorized,
    refresh,
    logout,
    onLogout,
  };
});
