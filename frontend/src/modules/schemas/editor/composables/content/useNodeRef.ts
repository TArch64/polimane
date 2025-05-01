import Konva from 'konva';
import { computed, type Ref, shallowRef } from 'vue';
import type { IKonvaNodeHolder } from 'vue-konva';

type MaybeNode = Konva.Node | null;
export type KonvaNodeRef<N extends MaybeNode> = Ref<N, IKonvaNodeHolder<N> | null>;

export function useNodeRef<N extends MaybeNode>(): KonvaNodeRef<N> {
  const nodeRef = shallowRef<N>(null!);

  return computed<N, IKonvaNodeHolder<N> | null>({
    get: () => nodeRef.value,
    set: (value) => nodeRef.value = value?.getNode() ?? null,
  });
}
