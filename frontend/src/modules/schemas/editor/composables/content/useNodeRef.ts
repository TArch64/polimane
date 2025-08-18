import Konva from 'konva';
import { computed, type Ref, shallowRef } from 'vue';
import type { IKonvaNodeHolder } from 'vue-konva';

type MaybeNode = Konva.Node | null;
export type KonvaNodeRef<N extends MaybeNode = MaybeNode> = Ref<N, IKonvaNodeHolder<Exclude<N, null>> | null>;

export function useNodeRef<N extends MaybeNode>(): KonvaNodeRef<N> {
  const nodeRef = shallowRef<N | null>(null);

  return computed<N, IKonvaNodeHolder<Exclude<N, null>> | null>({
    get: () => nodeRef.value!,
    set: (value) => nodeRef.value = value?.getNode() ?? null,
  });
}
