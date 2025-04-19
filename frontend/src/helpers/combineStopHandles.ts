import type { WatchStopHandle } from 'vue';

export function combineStopHandles(...handlers: WatchStopHandle[]): WatchStopHandle {
  return () => handlers.forEach((handle) => handle());
}
