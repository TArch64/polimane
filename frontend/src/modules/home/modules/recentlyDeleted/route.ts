import { defineViewRoute, type InferViewRouteInfo } from '@/router/define';

export const homeRecentlyDeletedRoute = defineViewRoute({
  name: 'home-recently-deleted',
  path: 'recently-deleted',
  component: () => import('./Route.vue'),
});

export type HomeRecentlyDeletedRouteInfo = InferViewRouteInfo<typeof homeRecentlyDeletedRoute>;
