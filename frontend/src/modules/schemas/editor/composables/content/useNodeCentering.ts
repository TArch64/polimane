import Konva from 'konva';
import { computed, type MaybeRefOrGetter, toValue } from 'vue';
import { reactiveComputed } from '@vueuse/core';
import { useNodeClientRect } from './useNodeClientRect';
import { useNodeParent } from './useNodeParent';

export type NodeCenteringPadding = Record<'vertical' | 'horizontal', number>;

export interface INodeCenteringOptions {
  padding?: NodeCenteringPadding;
}

export function useNodeCentering(nodeRef: MaybeRefOrGetter<Konva.Node | null>, options: INodeCenteringOptions = {}): Partial<Konva.NodeConfig> {
  const padding = computed(() => options.padding ?? { vertical: 0, horizontal: 0 });

  const node = computed(() => toValue(nodeRef));
  const nodeRect = useNodeClientRect(node);

  const parentNode = useNodeParent(node);

  return reactiveComputed((): Partial<Konva.NodeConfig> => {
    if (nodeRect.value?.isBlank || !parentNode.value) {
      return {};
    }

    const freeSpaceX = parentNode.value.width() - nodeRect.value.width - padding.value.horizontal * 2;
    const freeSpaceY = parentNode.value.height() - nodeRect.value.height - padding.value.vertical * 2;

    return {
      x: Math.max(freeSpaceX / 2, 0) + padding.value.horizontal,
      y: Math.max(freeSpaceY / 2, 0) + padding.value.vertical,
    };
  });
}
