import { defineViewRoute, type InferViewRouteInfo } from '@/router/define';

export const settingsProfileRoute = defineViewRoute({
  name: 'settings-profile',
  path: 'profile',
  component: () => import('./Route.vue'),
});

export type SettingsProfileRouteInfo = InferViewRouteInfo<typeof settingsProfileRoute>;
