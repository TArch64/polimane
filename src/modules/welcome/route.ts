import { defineRoute } from '@/router/define';

export const welcomeRoute = defineRoute({
  name: 'welcome',
  path: '/welcome',
  component: () => import('./Route.vue'),
});
