import {
  defineRedirectRoute,
  defineWrapperRoute,
  type InferWrapperRouteInfo,
} from '@/router/define';
import { settingsProfileRoute, type SettingsProfileRouteInfo } from './profile';

const notFoundRoute = defineRedirectRoute('', settingsProfileRoute.name);

export const settingsRoute = defineWrapperRoute({
  name: 'settings',
  path: '/settings',

  children: [
    settingsProfileRoute,
    notFoundRoute,
  ],
});

export type SettingsRouteInfo = InferWrapperRouteInfo<typeof settingsRoute, SettingsProfileRouteInfo>;
