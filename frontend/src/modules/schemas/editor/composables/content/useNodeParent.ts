import { computed, type ComputedRef, type MaybeRefOrGetter, toValue } from 'vue';
import Konva from 'konva';

export function useNodeParent(nodeRef: MaybeRefOrGetter<Konva.Node | null>): ComputedRef<Konva.Node | null> {
  const node = computed(() => toValue(nodeRef));
  return computed(() => node.value?.parent ?? null);
}
