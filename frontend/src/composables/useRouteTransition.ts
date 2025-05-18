import { nextTick } from 'vue';

export interface IRouteTransition {
  start: (callback: ViewTransitionUpdateCallback) => void;
}

let pending: ViewTransitionUpdateCallback[] = [];

export function useRouteTransition(): IRouteTransition {
  function doTransition(): void {
    document.startViewTransition(async () => {
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
