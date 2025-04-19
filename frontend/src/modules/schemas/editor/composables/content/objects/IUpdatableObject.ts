export interface IUpdatableObject<V> {
  update(value: V): void;
}

export function isUpdatableObject<V>(object: unknown): object is IUpdatableObject<V> {
  return typeof (object as IUpdatableObject<V>).update === 'function';
}
