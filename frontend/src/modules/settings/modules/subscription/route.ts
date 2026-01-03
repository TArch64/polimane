import { defineViewRoute, type InferViewRouteInfo } from '@/router/define';

export const settingsSubscriptionRoute = defineViewRoute({
  name: 'settings-subscription',
  path: 'subscription',
  component: () => import('./Route.vue'),
});

export type SettingsSubscriptionRouteInfo = InferViewRouteInfo<typeof settingsSubscriptionRoute>;
