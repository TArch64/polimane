export function getObjectKeys<T extends object>(obj: T): (keyof T)[] {
  return Object.keys(obj) as (keyof T)[];
}

export type ObjectEntries<T extends object> = {
  [K in keyof T]: [K, T[K]];
}[keyof T][];

export function getObjectEntries<T extends object>(obj: T): ObjectEntries<T> {
  return Object.entries(obj) as ObjectEntries<T>;
}
