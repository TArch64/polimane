import {
  computed,
  type MaybeRefOrGetter,
  onBeforeUnmount,
  type ShallowRef,
  shallowRef,
  toValue,
  watch,
} from 'vue';
import Konva from 'konva';
import { useDebounceFn } from '@vueuse/core';
import { NodeRect } from '@/models';

export function useNodeClientRect(node: MaybeRefOrGetter<Konva.Node | null>): ShallowRef<NodeRect> {
  const nodeRef = computed(() => toValue(node));
  const clientRect = shallowRef(NodeRect.BLANK);

  const update = useDebounceFn(() => {
    if (!nodeRef.value) {
      clientRect.value = NodeRect.BLANK;
      return;
    }

    const newRect = nodeRef.value!.getClientRect();

    if (!clientRect.value.isEqual(newRect)) {
      clientRect.value = new NodeRect(newRect);
    }
  }, 10);

  watch(nodeRef, (node, oldNode) => {
    if (oldNode) oldNode.off('layout', update);

    if (node) {
      node.on('layout', update);
    }

    update();
  }, { immediate: true });

  onBeforeUnmount(() => nodeRef.value?.off('layout', update));

  return clientRect;
}
