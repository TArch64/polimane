export interface IDestroyableObject {
  destroy(): void;
}

export function isDestroyableObject(object: unknown): object is IDestroyableObject {
  return typeof (object as IDestroyableObject).destroy === 'function';
}
