import { defineWrapperRoute, type InferWrapperRouteInfo } from '@/router/define';
import { homeListRoute, type HomeListRouteInfo } from './modules/list';
import { homeFolderRoute, type HomeFolderRouteInfo } from './modules/folder';

export const homeRoute = defineWrapperRoute({
  path: '/',
  component: () => import('./Route.vue'),
  children: [homeListRoute, homeFolderRoute],
});

export type HomeRouteInfo = InferWrapperRouteInfo<typeof homeRoute, HomeListRouteInfo & HomeFolderRouteInfo>;
