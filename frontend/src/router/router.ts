import { createRouter, createWebHistory } from 'vue-router';
import { nextTick } from 'vue';
import { sessionMiddleware } from '@/router/middleware';
import { homeListRoute, homeRoute } from '@/modules/home';
import { schemasRoute } from '@/modules/schemas';
import { authRoute } from '@/modules/auth';
import { settingsRoute } from '@/modules/settings';
import { defineRedirectRoute } from '@/router/define';
import { startViewTransition } from '@/helpers';

const routes = [
  authRoute,
  homeRoute,
  schemasRoute,
  settingsRoute,
];

const notFoundRoute = defineRedirectRoute('/:pathMatch(.*)*', homeListRoute.name);

export const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),

  routes: [
    ...routes,
    notFoundRoute,
  ],
});

router.beforeEach(sessionMiddleware);

router.beforeResolve(async (_, __, next) => {
  const transition = startViewTransition(async () => {
    next();
    await nextTick();
    await nextTick();
  });

  await transition.ready;
});

router.onError((error) => {
  if (error.message.includes('Failed to fetch dynamically imported module')) {
    const reloadKey = 'router-reload-attempted';

    if (sessionStorage.getItem(reloadKey)) {
      sessionStorage.removeItem(reloadKey);
      throw new Error('Failed to load module after reload:', error);
    }

    sessionStorage.setItem(reloadKey, '1');
    window.location.reload();
  }
});
