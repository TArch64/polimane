import { createRouter, createWebHistory } from 'vue-router';
import { homeRoute } from '../modules/home';

export const routes = [homeRoute] as const;

export const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});
