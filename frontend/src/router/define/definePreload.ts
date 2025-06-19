import type {
  NavigationGuard,
  RouteLocationNormalizedTyped,
  RouteLocationRaw,
  RouteMap,
} from 'vue-router';

type RouteName = keyof RouteMap;
type RouteLocation<N extends RouteName> = RouteLocationNormalizedTyped<RouteMap, N>;
export type RoutePreloadResult = void | undefined | RouteLocationRaw;
export type RoutePreload<N extends RouteName> = (route: RouteLocation<N>) => Promise<RoutePreloadResult>;

export function definePreload<N extends RouteName = RouteName>(preload: RoutePreload<N>): NavigationGuard {
  return async (to, _, next) => {
    try {
      const result = await preload(to as RouteLocation<N>);
      result === undefined ? next() : next(result);
    } catch (error) {
      return next(error as Error);
    }
  };
}
