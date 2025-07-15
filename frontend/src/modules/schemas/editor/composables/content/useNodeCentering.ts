import Konva from 'konva';
import { type MaybeRefOrGetter, nextTick, ref, watch } from 'vue';
import { toReactive, toRef } from '@vueuse/core';
import { useNodeClientRect } from './useNodeClientRect';
import { useNodeParent } from './useNodeParent';
import { type NodePaddingInput, useNodePadding } from './useNodePadding';

export interface INodeCenteringOptions {
  padding?: NodePaddingInput;
}

export function useNodeCentering(nodeRef: MaybeRefOrGetter<Konva.Node | null>, options: INodeCenteringOptions = {}): Partial<Konva.NodeConfig> {
  const node = toRef(nodeRef);
  const nodeRect = useNodeClientRect(node);
  const parentNode = useNodeParent(node);
  const padding = useNodePadding(options.padding ?? 0);
  const config = ref<Partial<Konva.NodeConfig>>({});

  watch(nodeRect, async () => {
    await nextTick();

    if (nodeRect.value?.isBlank || !parentNode.value) {
      config.value = {};
      return;
    }

    const freeSpaceX = parentNode.value.width() - nodeRect.value.width - padding.horizontal * 2;
    const freeSpaceY = parentNode.value.height() - nodeRect.value.height - padding.vertical * 2;

    config.value = {
      x: Math.max(freeSpaceX / 2, 0) + padding.horizontal,
      y: Math.max(freeSpaceY / 2, 0) + padding.vertical,
    };
  }, { immediate: true });

  return toReactive(config);
}
