import { computed, type ComputedRef, type MaybeRefOrGetter, toValue } from 'vue';
import Konva from 'konva';

export function useNodeParent(node: MaybeRefOrGetter<Konva.Node | null>): ComputedRef<Konva.Node | null> {
  const nodeRef = computed(() => toValue(node));
  return computed(() => nodeRef.value?.parent ?? null);
}
