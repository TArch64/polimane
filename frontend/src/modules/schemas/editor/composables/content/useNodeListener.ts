import { type MaybeRefOrGetter, onBeforeUnmount, watch } from 'vue';
import { toRef } from '@vueuse/core';
import type { KonvaEventListener, Node, NodeEventMap } from 'konva/lib/Node';

export function useNodeListener<N extends Node | null, E extends keyof NodeEventMap>(
  nodeRef: MaybeRefOrGetter<N>,
  event: E,
  callback: KonvaEventListener<N, NodeEventMap[E]>,
) {
  const node = toRef(nodeRef);

  watch(node, (node, oldNode) => {
    if (oldNode) oldNode.off(event, callback);

    if (node) {
      // @ts-expect-error hard to type correctly
      node.on(event, callback);
    }
  }, { immediate: true });

  onBeforeUnmount(() => {
    node.value?.off(event, callback);
  });
}
