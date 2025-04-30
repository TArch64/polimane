import Konva from 'konva';
import type { KonvaComponent } from 'vue-konva';
import { computed, type Ref, shallowRef } from 'vue';
import type { SafeAny } from '@/types';

export function useNodeRef<N extends Konva.Node | null>(): Ref<N, SafeAny | null> {
  const nodeRef = shallowRef<N>(null!);

  return computed<N, KonvaComponent<N> | null>({
    get: () => nodeRef.value,
    set: (value) => nodeRef.value = value?.getNode() ?? null,
  });
}
