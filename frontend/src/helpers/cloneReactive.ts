export function cloneReactiveAsRaw<T>(obj: T): T {
  return JSON.parse(JSON.stringify(obj));
}
