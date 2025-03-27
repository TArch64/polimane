import { createRouter, createWebHistory, type NavigationGuardReturn } from 'vue-router';
import { useStorageStore } from '../stores';
import { homeRoute } from '../modules/home';
import { welcomeRoute } from '../modules/welcome';

export const routes = [welcomeRoute, homeRoute] as const;

export const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),

  routes: [
    {
      path: '/',
      children: [...routes],

      async beforeEnter(to): Promise<NavigationGuardReturn> {
        if (to.name === 'welcome') {
          return;
        }

        const storageStore = useStorageStore();
        await storageStore.loadState();

        if (!storageStore.isDirectorySelected) {
          return { name: 'welcome' };
        }
      },
    },
  ],
});
