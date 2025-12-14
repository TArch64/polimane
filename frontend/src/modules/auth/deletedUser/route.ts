import { defineViewRoute, type InferViewRouteInfo } from '@/router/define';

export const authDeletedUserRoute = defineViewRoute({
  name: 'auth-deleted-user',
  path: 'deleted',
  component: () => import('./Route.vue'),
});

export type AuthDeletedUserRouteInfo = InferViewRouteInfo<typeof authDeletedUserRoute>;
