import type { NavigationGuardWithThis } from 'vue-router';
import { useSessionStore } from '@/stores';

export const sessionMiddleware: NavigationGuardWithThis<undefined> = async (to, from, next) => {
  const sessionStore = useSessionStore();

  if (to.name === 'welcome') {
    return next();
  }

  await sessionStore.refresh();

  if (sessionStore.isLoginedIn) {
    return next();
  }

  return next({ name: 'welcome' });
};
