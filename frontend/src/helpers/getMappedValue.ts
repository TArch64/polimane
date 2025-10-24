import { type MaybeRefOrGetter, toValue } from 'vue';

export function getMappedValue<V extends string | number, M>(value: V, map: Record<V, MaybeRefOrGetter<M>>): M {
  return toValue(map[value]) as M;
}
