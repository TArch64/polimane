import { createRouter, createWebHistory } from 'vue-router';
import { sessionMiddleware } from '@/router/middleware';
import { homeRoute } from '@/modules/home';
import { welcomeRoute } from '@/modules/welcome';
import { schemasRoute } from '@/modules/schemas';

const routes = [
  welcomeRoute,
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
