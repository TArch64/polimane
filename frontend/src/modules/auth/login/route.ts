import { defineViewRoute, type InferViewRouteInfo } from '@/router/define';

export const authLoginRoute = defineViewRoute({
  name: 'auth',
  path: '',
  component: () => import('./Route.vue'),
});

export type AuthLoginRouteInfo = InferViewRouteInfo<typeof authLoginRoute>;
