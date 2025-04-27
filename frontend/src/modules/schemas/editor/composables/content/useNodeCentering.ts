import Konva from 'konva';
import { computed, type MaybeRefOrGetter, onMounted, toValue, watch } from 'vue';
import { useDebounceFn } from '@vueuse/core';

export type NodeCenteringPadding = Record<'vertical' | 'horizontal', number>;

export interface INodeCenteringOptions {
  padding?: NodeCenteringPadding;
  trigger?: MaybeRefOrGetter<unknown>;
}

export interface INodeCentering {
  update: () => void;
}

export function useNodeCentering(nodeRef: MaybeRefOrGetter<Konva.Node | null>, options: INodeCenteringOptions = {}): INodeCentering {
  const node = computed(() => toValue(nodeRef));
  const trigger = options.trigger ? computed(() => toValue(options.trigger)) : null;
  const padding = computed(() => options.padding ?? { vertical: 0, horizontal: 0 });

  const update = useDebounceFn(() => {
    const parent = node.value?.parent;

    if (!node.value || !parent) {
      return;
    }

    const nodeRect = node.value.getClientRect();
    const freeSpaceX = parent.width() - nodeRect.width - padding.value.horizontal * 2;
    const freeSpaceY = parent.height() - nodeRect.height - padding.value.vertical * 2;

    node.value.position({
      x: Math.max(freeSpaceX / 2, 0) + padding.value.horizontal,
      y: Math.max(freeSpaceY / 2, 0) + padding.value.vertical,
    });
  }, 30);

  onMounted(update);

  if (trigger) {
    watch(trigger, update, { deep: true });
  }

  return { update };
}
