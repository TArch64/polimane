import type { HomeRouteInfo } from '@/modules/home';
import type { SchemasRouteInfo } from '@/modules/schemas';
import type { AuthRouteInfo } from '@/modules/auth/route';

type RouteNamedMap = AuthRouteInfo & HomeRouteInfo & SchemasRouteInfo;

declare module 'vue-router' {
  interface TypesConfig {
    RouteNamedMap: RouteNamedMap;
  }
}

export {};
