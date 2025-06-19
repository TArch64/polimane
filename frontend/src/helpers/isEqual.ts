import { equals } from 'rambda';

export function isEqual(v1: unknown, v2: unknown): boolean {
  return equals(v1, v2);
}
