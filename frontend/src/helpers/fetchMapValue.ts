export function fetchMapValue<K, V>(map: Map<K, V>, key: K, fetch: () => V): V {
  if (!map.has(key)) map.set(key, fetch());
  return map.get(key)!;
}
