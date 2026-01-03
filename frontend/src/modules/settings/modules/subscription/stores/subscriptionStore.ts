import { defineStore } from 'pinia';
import { computed } from 'vue';
import { useSessionStore } from '@/stores';

export const useSubscriptionStore = defineStore('settings/subscription', () => {
  const sessionStore = useSessionStore();
  const subscription = computed(() => sessionStore.user.subscription);

  return { subscription };
});
