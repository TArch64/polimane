import { defineRoute, type InferViewRouteInfo } from '@/router/define';

export const authRoute = defineRoute({
  name: 'auth',
  path: '/auth',
  component: () => import('./Route.vue'),
});

export type AuthRouteInfo = InferViewRouteInfo<typeof authRoute>;
