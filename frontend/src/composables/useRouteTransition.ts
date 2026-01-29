import { nextTick } from 'vue';
import { startViewTransition, type ViewTransitionState } from '@/helpers';

export interface IRouteTransition extends ViewTransitionState {
  start: (callback: ViewTransitionUpdateCallback) => void;
}

let pending: ViewTransitionUpdateCallback[] = [];
let readyResolver: PromiseWithResolvers<void>;
let finishedResolver: PromiseWithResolvers<void>;

export function useRouteTransition(): IRouteTransition {
  function doTransition(): void {
    const transition = startViewTransition(async () => {
      for (const callback of pending) {
        await callback();
      }
      pending = [];
    });

    transition.ready.then(() => readyResolver.resolve());
    transition.finished.then(() => finishedResolver.resolve());
  }

  function start(callback: ViewTransitionUpdateCallback): void {
    pending.push(callback);

    if (pending.length === 1) {
      finishedResolver = Promise.withResolvers();
      readyResolver = Promise.withResolvers();
      nextTick().then(doTransition);
    }
  }

  return {
    start,

    get ready() {
      return readyResolver.promise;
    },

    get finished() {
      return finishedResolver.promise;
    },
  };
}
