import type {
  NavigationGuard,
  RouteLocationNormalizedTyped,
  RouteLocationRaw,
  RouteMap,
} from 'vue-router';
import { useDebounceFn } from '@vueuse/core';
import { done as doneProgress, start as startProgress } from 'nprogress';

type RouteName = keyof RouteMap;
type RouteLocation<N extends RouteName> = RouteLocationNormalizedTyped<RouteMap, N>;
export type RoutePreloadResult = void | undefined | RouteLocationRaw;
export type RoutePreload<N extends RouteName> = (route: RouteLocation<N>) => Promise<RoutePreloadResult>;

type ProgressState = 'start' | 'success' | 'error';

export function definePreload<N extends RouteName = RouteName>(preload: RoutePreload<N>): NavigationGuard {
  const toggleProgress = useDebounceFn((state: ProgressState) => {
    state === 'start' ? startProgress() : doneProgress(state === 'error');
  }, 100);

  return async (to, _, next) => {
    try {
      toggleProgress('start');
      const result = await preload(to as RouteLocation<N>);
      await toggleProgress('success');
      result === undefined ? next() : next(result);
    } catch (error) {
      await toggleProgress('error');
      return next(error as Error);
    }
  };
}
