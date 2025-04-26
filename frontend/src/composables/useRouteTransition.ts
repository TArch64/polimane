import { nextTick } from 'vue';

export interface IRouteTransition {
  start: (callback: ViewTransitionUpdateCallback) => void;
}

export function useRouteTransition(): IRouteTransition {
  let pending: ViewTransitionUpdateCallback[] = [];
  let isScheduled = false;

  function doTransition(): void {
    document.startViewTransition(async () => {
      for (const callback of pending) {
        await callback();
      }
      pending = [];
      isScheduled = false;
    });
  }

  function start(callback: ViewTransitionUpdateCallback): void {
    pending.push(callback);

    if (!isScheduled) {
      isScheduled = true;
      nextTick().then(doTransition);
    }
  }

  return { start };
}
