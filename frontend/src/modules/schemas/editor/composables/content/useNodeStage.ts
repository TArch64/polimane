import { computed, type MaybeRefOrGetter, toValue } from 'vue';
import Konva from 'konva';

export function useNodeStage(nodeRef: MaybeRefOrGetter<Konva.Node | null>) {
  return computed(() => toValue(nodeRef)?.getStage() ?? null);
}
