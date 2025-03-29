import { createRouter, createWebHistory } from 'vue-router';
import { homeRoute } from '@/modules/home';
import { welcomeRoute } from '@/modules/welcome';
import { checkStorageDirMiddleware } from './middleware';

export const routes = [
  welcomeRoute,
  homeRoute,
] as const;

export const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),

  routes: [
    {
      path: '/',
      children: [...routes],
      beforeEnter: checkStorageDirMiddleware,
    },
  ],
});
