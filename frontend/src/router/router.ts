import { createRouter, createWebHistory } from 'vue-router';
import { homeRoute } from '@/modules/home';
import { welcomeRoute } from '@/modules/welcome';
import { sessionMiddleware } from '@/router/middleware';

export const routes = [
  welcomeRoute,
  homeRoute,
] as const;

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
