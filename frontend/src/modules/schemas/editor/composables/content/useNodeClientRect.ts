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

export function useNodeClientRect(nodeRef: MaybeRefOrGetter<Konva.Node | null>): ShallowRef<NodeRect> {
  const node = computed(() => toValue(nodeRef));
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

  watch(node, (node, oldNode) => {
    if (oldNode) oldNode.off('layout', update);

    if (node) {
      node.on('layout', update);
    }

    update();
  }, { immediate: true });

  onBeforeUnmount(() => node.value?.off('layout', update));

  return clientRect;
}
