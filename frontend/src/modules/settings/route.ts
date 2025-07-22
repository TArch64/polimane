import {
  defineRedirectRoute,
  defineWrapperRoute,
  type InferWrapperRouteInfo,
} from '@/router/define';
import { settingsProfileRoute, type SettingsProfileRouteInfo } from './profile';
import { settingsSecurityRoute, type SettingsSecurityRouteInfo } from './security';

const notFoundRoute = defineRedirectRoute('', settingsProfileRoute.name);

export const settingsRoute = defineWrapperRoute({
  path: '/settings',
  component: () => import('./Route.vue'),

  children: [
    settingsProfileRoute,
    settingsSecurityRoute,
    notFoundRoute,
  ],
});

export type SettingsRouteInfo = InferWrapperRouteInfo<typeof settingsRoute, SettingsProfileRouteInfo & SettingsSecurityRouteInfo>;
