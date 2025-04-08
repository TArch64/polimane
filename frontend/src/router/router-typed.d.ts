import type { HomeRouteInfo } from '@/modules/home';
import type { WelcomeRouteInfo } from '@/modules/welcome';
import type { SchemasRouteInfo } from '@/modules/schemas';

type RouteNamedMap = WelcomeRouteInfo & HomeRouteInfo & SchemasRouteInfo;

declare module 'vue-router' {
  interface TypesConfig {
    RouteNamedMap: RouteNamedMap;
  }
}

export {};
