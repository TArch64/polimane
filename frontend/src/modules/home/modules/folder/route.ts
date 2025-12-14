import { defineViewRoute, type InferViewRouteInfo } from '@/router/define';

export const homeFolderRoute = defineViewRoute({
  name: 'home-folder',
  path: 'folders/:folderId',
  component: () => import('./Route.vue'),
});

export type HomeFolderRouteInfo = InferViewRouteInfo<typeof homeFolderRoute>;
