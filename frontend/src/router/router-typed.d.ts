import type { HomeRouteInfo } from '@/modules/home';
import type { SchemasRouteInfo } from '@/modules/schemas';

type RouteNamedMap = HomeRouteInfo & SchemasRouteInfo;

declare module 'vue-router' {
  interface TypesConfig {
    RouteNamedMap: RouteNamedMap;
  }
}

export {};
