import type { NavigationGuard } from 'vue-router';
import { useSessionStore } from '@/stores';

export const sessionMiddleware: NavigationGuard = async (to, _, next) => {
  const sessionStore = useSessionStore();

  if (to.query['access-token'] && to.query['refresh-token']) {
    sessionStore.setTokens(
      to.query['access-token'] as string,
      to.query['refresh-token'] as string,
    );
  }

  if (!sessionStore.isLoggedIn) {
    await sessionStore.refresh();
  }

  sessionStore.isLoggedIn ? next() : next(false);
};
