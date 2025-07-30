import type { HomeRouteInfo } from '@/modules/home';
import type { SchemasRouteInfo } from '@/modules/schemas';
import type { AuthRouteInfo } from '@/modules/auth/route';
import type { SettingsRouteInfo } from '@/modules/settings';

type RouteNamedMap = AuthRouteInfo & HomeRouteInfo & SchemasRouteInfo & SettingsRouteInfo;

declare module 'vue-router' {
  interface TypesConfig {
    RouteNamedMap: RouteNamedMap;
  }
}

export {};
