import { createRouter, createWebHistory } from 'vue-router';
import { nextTick } from 'vue';
import { sessionMiddleware } from '@/router/middleware';
import { homeRoute } from '@/modules/home';
import { schemasRoute } from '@/modules/schemas';

const routes = [
  homeRoute,
  schemasRoute,
];

export const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),

  routes: [
    ...routes,
    {
      path: '/:pathMatch(.*)*',
      redirect: { name: 'home' },
    },
  ],
});

router.beforeEach(sessionMiddleware);

router.beforeResolve(async (_, __, next) => {
  const transition = document.startViewTransition(async () => {
    next();
    await nextTick();
    await nextTick();
  });

  await transition.ready;
});
