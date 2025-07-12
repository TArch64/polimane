if (typeof document.startViewTransition !== 'function') {
  document.startViewTransition = (callback: () => Promise<void>): ViewTransition => ({
    ready: callback(),
    skipTransition: null!,
    types: null!,
    finished: null!,
    updateCallbackDone: null!,
  });
}
