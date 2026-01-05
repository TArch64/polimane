import { defineViewRoute, type InferViewRouteInfo } from '@/router/define';

export const settingsSecurityRoute = defineViewRoute({
  name: 'settings-security',
  path: 'security',
  component: () => import('./Route.vue'),
});

export type SettingsSecurityRouteInfo = InferViewRouteInfo<typeof settingsSecurityRoute>;
