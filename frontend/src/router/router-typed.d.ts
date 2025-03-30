import type { RouteRecordInfo } from 'vue-router';
import { routes } from './router';

type AppRoutes = (typeof routes)[number];

type RouteNamedMap = {
  [R in AppRoutes as R['name']]: RouteRecordInfo<
    R['name'],
    R['path'],
    Record<never, never>,
    Record<never, never>
  >;
};

declare module 'vue-router' {
  interface TypesConfig {
    RouteNamedMap: RouteNamedMap;
  }
}

export {};
