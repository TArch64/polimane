import type { SafeAny } from './SafeAny';

export type Constructor<V> = {
  new(...args: SafeAny[]): V;
};
