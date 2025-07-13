import { defineRoute, type InferViewRouteInfo } from '@/router/define';

export const authCompleteRoute = defineRoute({
  name: 'authComplete',
  path: 'complete',
  component: () => import('./Route.vue'),
});

export type AuthCompleteRouteInfo = InferViewRouteInfo<typeof authCompleteRoute>;
