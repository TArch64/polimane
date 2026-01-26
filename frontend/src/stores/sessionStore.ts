import { defineStore } from 'pinia';
import { computed, nextTick, type Ref, ref, watch } from 'vue';
import type { IUser, IUserSubscription, UserActivePlan } from '@/models';
import {
  type HttpBody,
  UpdateCountersMiddleware,
  useAuthorized,
  useHttpClient,
} from '@/composables';
import { SubscriptionLimit, SubscriptionPlanId } from '@/enums';

interface IChangePlanBody {
  planId: SubscriptionPlanId;
}

export const useSessionStore = defineStore('session', () => {
  const http = useHttpClient();
  const updateCountersMiddleware = http.getMiddleware(UpdateCountersMiddleware)!;

  const authorized = useAuthorized();

  const user = ref<IUser | null>(null);
  const subscription = computed(() => user.value?.subscription);
  const plan = computed(() => subscription.value?.plan);

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

  const getLimit = (name: SubscriptionLimit) => plan.value?.limits[name];

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

  async function changePlan(planId: SubscriptionPlanId): Promise<void> {
    await http.post<HttpBody, IChangePlanBody>('/users/current/subscription/change', {
      planId,
    });

    try {
      await refresh();
    } catch {
      location.reload();
    }
  }

  return {
    user: user as Ref<IUser>,
    subscription: subscription as Ref<IUserSubscription>,
    plan: plan as Ref<UserActivePlan>,
    updateUser,
    getLimit,
    isLoggedIn: authorized,
    refresh,
    logout,
    onLogout,
    changePlan,
  };
});
