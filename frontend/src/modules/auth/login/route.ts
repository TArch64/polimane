import { defineRoute, type InferViewRouteInfo } from '@/router/define';

export const authLoginRoute = defineRoute({
  name: 'auth',
  path: '',
  component: () => import('./Route.vue'),
});

export type AuthLoginRouteInfo = InferViewRouteInfo<typeof authLoginRoute>;
