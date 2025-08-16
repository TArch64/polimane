import { toRef } from '@vueuse/core';
import { computed, type MaybeRefOrGetter, onBeforeUnmount, toValue, watch } from 'vue';
import type { KonvaEventListener, Node, NodeEventMap } from 'konva/lib/Node';

export interface INodeListenerOptions {
  isActive?: MaybeRefOrGetter<boolean>;
}

export function useNodeListener<N extends Node | null, E extends keyof NodeEventMap>(
  nodeRef: MaybeRefOrGetter<N>,
  event: E,
  callback: KonvaEventListener<N, NodeEventMap[E]>,
  options: INodeListenerOptions = {},
) {
  const node = toRef(nodeRef);
  const isActive = computed(() => toValue(options.isActive) ?? true);

  watch([node, isActive], ([node, isActive], [oldNode]) => {
    if (!isActive) {
      node?.off(event, callback);
      oldNode?.off(event, callback);
      return;
    }

    oldNode?.off(event, callback);
    // @ts-expect-error hard to type correctly
    node?.on(event, callback);
  }, { immediate: true });

  onBeforeUnmount(() => {
    node.value?.off(event, callback);
  });
}
