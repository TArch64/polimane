import { defineWrapperRoute, type InferWrapperRouteInfo } from '@/router/define';
import { homeListRoute, type HomeListRouteInfo } from './modules/list';
import { homeFolderRoute, type HomeFolderRouteInfo } from './modules/folder';
import {
  homeRecentlyDeletedRoute,
  type HomeRecentlyDeletedRouteInfo,
} from './modules/recentlyDeleted';

export const homeRoute = defineWrapperRoute({
  path: '/',
  component: () => import('./Route.vue'),
  children: [
    homeListRoute,
    homeRecentlyDeletedRoute,
    homeFolderRoute,
  ],
});

export type HomeRouteInfo = InferWrapperRouteInfo<typeof homeRoute, HomeListRouteInfo & HomeFolderRouteInfo & HomeRecentlyDeletedRouteInfo>;
