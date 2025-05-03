import Konva from 'konva';
import {
  computed,
  type MaybeRefOrGetter,
  onBeforeUnmount,
  onMounted,
  reactive,
  toValue,
} from 'vue';
import { useDebounceFn } from '@vueuse/core';

export type NodeCenteringPadding = Record<'vertical' | 'horizontal', number>;

export interface INodeCenteringOptions {
  padding?: NodeCenteringPadding;
}

export function useNodeCentering(nodeRef: MaybeRefOrGetter<Konva.Node | null>, options: INodeCenteringOptions = {}): Partial<Konva.NodeConfig> {
  const node = computed(() => toValue(nodeRef));
  const padding = computed(() => options.padding ?? { vertical: 0, horizontal: 0 });
  const config = reactive<Partial<Konva.NodeConfig>>({});

  async function update(event?: unknown) {
    console.log(event);

    const parent = node.value?.parent;

    if (!node.value || !parent) {
      return;
    }

    const nodeRect = node.value.getClientRect();
    const freeSpaceX = parent.width() - nodeRect.width - padding.value.horizontal * 2;
    const freeSpaceY = parent.height() - nodeRect.height - padding.value.vertical * 2;

    config.x = Math.max(freeSpaceX / 2, 0) + padding.value.horizontal;
    config.y = Math.max(freeSpaceY / 2, 0) + padding.value.vertical;
  }

  const scheduleUpdate = useDebounceFn(update, 10);

  onMounted(() => {
    update();
    node.value?.on('layout', scheduleUpdate);
  });

  onBeforeUnmount(() => {
    node.value?.off('layout', update);
  });

  return config;
}
