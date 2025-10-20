import { nextTick } from 'vue';
import { startViewTransition } from '@/helpers';

export interface IRouteTransition {
  start: (callback: ViewTransitionUpdateCallback) => void;
}

let pending: ViewTransitionUpdateCallback[] = [];

export function useRouteTransition(): IRouteTransition {
  function doTransition(): void {
    startViewTransition(async () => {
      for (const callback of pending) {
        await callback();
      }
      pending = [];
    });
  }

  function start(callback: ViewTransitionUpdateCallback): void {
    pending.push(callback);

    if (pending.length === 1) {
      nextTick().then(doTransition);
    }
  }

  return { start };
}
