export function combineStopHandles(...handlers: VoidFunction[]): VoidFunction {
  return () => handlers.forEach((handle) => handle());
}
