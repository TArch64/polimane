import { type MaybeRefOrGetter, onMounted, type ShallowRef, shallowRef } from 'vue';
import Konva from 'konva';
import { toRef, useDebounceFn } from '@vueuse/core';
import { NodeRect } from '@/models';
import { useNodeListener } from './useNodeListener';

export function useNodeClientRect(nodeRef: MaybeRefOrGetter<Konva.Node | null>): ShallowRef<NodeRect> {
  const node = toRef(nodeRef);
  const clientRect = shallowRef(NodeRect.BLANK);

  const update = useDebounceFn(() => {
    if (!node.value) {
      clientRect.value = NodeRect.BLANK;
      return;
    }

    const newRect = node.value!.getClientRect();

    if (!clientRect.value.isEqual(newRect)) {
      clientRect.value = new NodeRect(newRect);
    }
  }, 10);

  onMounted(update);
  useNodeListener(nodeRef, 'layout', update);

  return clientRect;
}
