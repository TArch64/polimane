import type { NavigationGuard } from 'vue-router';
import { useSessionStore } from '@/stores';

export const sessionMiddleware: NavigationGuard = async (to, _, next) => {
  const sessionStore = useSessionStore();

  if (to.name === 'welcome') {
    return next();
  }

  if (!sessionStore.isLoginedIn) {
    await sessionStore.refresh();
  }

  if (sessionStore.isLoginedIn) {
    return next();
  }

  return next({ name: 'welcome' });
};
