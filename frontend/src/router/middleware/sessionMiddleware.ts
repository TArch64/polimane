import type { NavigationGuard } from 'vue-router';
import { useSessionStore } from '@/stores';

export const sessionMiddleware: NavigationGuard = async (to, _, next) => {
  const sessionStore = useSessionStore();

  if (to.name === 'auth') {
    return next();
  }

  if (!sessionStore.isLoggedIn) {
    await sessionStore.refresh();
  }

  if (sessionStore.isLoggedIn) {
    return next();
  }

  return next({
    name: 'auth',
    query: { 'return-to': to.fullPath },
  });
};
