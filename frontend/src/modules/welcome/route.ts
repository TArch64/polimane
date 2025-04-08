import { defineRoute, type InferViewRouteInfo } from '@/router/define';

export const welcomeRoute = defineRoute({
  name: 'welcome',
  path: '/welcome',
  component: () => import('./Route.vue'),
});

export type WelcomeRouteInfo = InferViewRouteInfo<typeof welcomeRoute>;
