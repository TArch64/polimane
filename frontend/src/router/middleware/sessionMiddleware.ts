import type { NavigationGuard } from 'vue-router';
import { useSessionStore } from '@/stores';

export const sessionMiddleware: NavigationGuard = async (to, from, next) => {
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
