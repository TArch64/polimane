export function getObjectKeys<T extends object>(obj: T): (keyof T)[] {
  return Object.keys(obj) as (keyof T)[];
}

export type ObjectEntries<T extends object> = {
  [K in keyof T]: [K, T[K]];
}[keyof T][];

export type ObjectEntry<T extends object> = {
  [K in keyof T]: [K, T[K]];
}[keyof T];

export function getObjectEntries<T extends object>(obj: T): ObjectEntries<T> {
  return Object.entries(obj) as ObjectEntries<T>;
}

export function deleteObjectKeys<T extends object>(obj: T, keys: Set<keyof T>): T {
  const current = getObjectEntries(obj);
  const updated = current.filter(([key]) => !keys.has(key));
  return Object.fromEntries(updated) as T;
}
