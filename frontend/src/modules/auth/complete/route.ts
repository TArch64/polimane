import { defineViewRoute, type InferViewRouteInfo } from '@/router/define';

export const authCompleteRoute = defineViewRoute({
  name: 'auth-complete',
  path: 'complete',
  component: () => import('./Route.vue'),
});

export type AuthCompleteRouteInfo = InferViewRouteInfo<typeof authCompleteRoute>;
