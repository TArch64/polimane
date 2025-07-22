import {
  defineRedirectRoute,
  defineWrapperRoute,
  type InferWrapperRouteInfo,
} from '@/router/define';
import { settingsProfileRoute, type SettingsProfileRouteInfo } from './profile';

const notFoundRoute = defineRedirectRoute('', settingsProfileRoute.name);

export const settingsRoute = defineWrapperRoute({
  path: '/settings',
  component: () => import('./Route.vue'),

  children: [
    settingsProfileRoute,
    notFoundRoute,
  ],
});

export type SettingsRouteInfo = InferWrapperRouteInfo<typeof settingsRoute, SettingsProfileRouteInfo>;
