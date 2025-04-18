import { type MaybeRefOrGetter, toValue } from 'vue';

export function getMapValue<V extends string, M>(value: V, map: Record<V, MaybeRefOrGetter<M>>): M {
  return toValue(map[value]) as M;
}
