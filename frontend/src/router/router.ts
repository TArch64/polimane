import { createRouter, createWebHistory } from 'vue-router';
import { nextTick } from 'vue';
import { homeRoute } from '@/modules/home';
import { schemasRoute } from '@/modules/schemas';
import { authRoute } from '@/modules/auth';
import { settingsRoute } from '@/modules/settings';
import { startViewTransition } from '@/helpers';
import { defineRedirectRoute } from './define';
import { googleAnalytics, sessionMiddleware } from './middleware';

const routes = [
  authRoute,
  homeRoute,
  schemasRoute,
  settingsRoute,
];

const notFoundRoute = defineRedirectRoute('/:pathMatch(.*)*', homeRoute.name);

export const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),

  routes: [
    ...routes,
    notFoundRoute,
  ],
});

router.beforeEach(sessionMiddleware);
router.afterEach(googleAnalytics);

router.beforeResolve(async (_, __, next) => {
  const transition = startViewTransition(async () => {
    next();
    await nextTick();
    await nextTick();
  });

  await transition.ready;
});
