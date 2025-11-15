import { defineWrapperRoute, type InferWrapperRouteInfo } from '@/router/define';
import { homeListRoute, type HomeListRouteInfo } from './modules/list';

export const homeRoute = defineWrapperRoute({
  path: '/',
  component: () => import('./Route.vue'),
  children: [homeListRoute],
});

export type HomeRouteInfo = InferWrapperRouteInfo<typeof homeRoute, HomeListRouteInfo>;
