import type {
  NavigationGuard,
  RouteLocationNormalizedTyped,
  RouteLocationRaw,
  RouteMap,
} from 'vue-router';
import { useDebounceFn } from '@vueuse/core';
import { done as doneProgress, start as startProgress } from 'nprogress';
import { type ILoader, useLoaderStore } from '@/stores';

type RouteName = keyof RouteMap;
type RouteLocation<N extends RouteName> = RouteLocationNormalizedTyped<RouteMap, N>;
export type RoutePreloadResult = void | undefined | RouteLocationRaw;
export type RoutePreload<N extends RouteName> = (route: RouteLocation<N>) => Promise<RoutePreloadResult>;

const nProgressLoader: ILoader = {
  show: startProgress,
  hide: doneProgress,
};

export interface IPreloadOptions {
  appLoader?: boolean;
}

export function definePreload<N extends RouteName = RouteName>(preload: RoutePreload<N>, options: IPreloadOptions = {}): NavigationGuard {
  const loaderStore = useLoaderStore();
  const loader = options.appLoader ? loaderStore : nProgressLoader;

  const toggleProgress = useDebounceFn((isActive: boolean) => {
    isActive ? loader.show() : loader.hide();
  }, 100);

  return async (to, _, next) => {
    try {
      if (options.appLoader) {
        loaderStore.show();
      }

      toggleProgress(true);
      const result = await preload(to as RouteLocation<N>);
      result === undefined ? next() : next(result);
    } catch (error) {
      return next(error as Error);
    } finally {
      await toggleProgress(false);
      loaderStore.hide();
    }
  };
}
