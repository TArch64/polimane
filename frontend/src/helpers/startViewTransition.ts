const fallback = (callback: () => Promise<void>) => ({
  ready: callback(),
  skipTransition: null!,
  types: null!,
  finished: null!,
  updateCallbackDone: null!,
});

export function startViewTransition(callback: () => Promise<void>): ViewTransition {
  const canExecute = typeof document.startViewTransition === 'function'
    && document.readyState === 'complete'
    && document.visibilityState === 'visible';

  return canExecute ? document.startViewTransition(callback) : fallback(callback);
}
