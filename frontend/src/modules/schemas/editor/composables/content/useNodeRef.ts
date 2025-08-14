import Konva from 'konva';
import { computed, type Ref, type ShallowRef, shallowRef } from 'vue';
import type { IKonvaNodeHolder } from '@/modules/schemas/editor/components/content/konva';

type MaybeNode = Konva.Node | null;
export type KonvaNodeRef<N extends MaybeNode = MaybeNode> = Ref<N, IKonvaNodeHolder<Exclude<N, null>> | null>;

export function useNodeRef<N extends MaybeNode>(storeRef?: ShallowRef<N | null>): KonvaNodeRef<N> {
  const nodeRef = storeRef ?? shallowRef<N | null>(null);

  return computed<N, IKonvaNodeHolder<Exclude<N, null>> | null>({
    get: () => nodeRef.value!,
    set: (value) => nodeRef.value = value?.getNode() ?? null,
  });
}
