import { nextTick } from 'vue';
import { startViewTransition } from '@/helpers';

export interface IRouteTransition {
  start: (callback: ViewTransitionUpdateCallback) => Promise<void>;
}

let pending: ViewTransitionUpdateCallback[] = [];

export function useRouteTransition(): IRouteTransition {
  let resolvers: PromiseWithResolvers<void>;

  function doTransition(): void {
    startViewTransition(async () => {
      for (const callback of pending) {
        await callback();
      }
      pending = [];
      resolvers.resolve();
    });
  }

  function start(callback: ViewTransitionUpdateCallback): Promise<void> {
    pending.push(callback);

    if (pending.length === 1) {
      resolvers = Promise.withResolvers();
      nextTick().then(doTransition);
    }

    return resolvers.promise;
  }

  return { start };
}
