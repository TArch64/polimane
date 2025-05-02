import { computed, type MaybeRefOrGetter, reactive, ref, toValue } from 'vue';
import type { KonvaEventObject } from 'vue-konva';
import { whenever } from '@vueuse/core';

export interface INodeHoverOptions {
  isDisabled?: MaybeRefOrGetter<boolean>;
}

type NodeHoverListener = (event: KonvaEventObject<MouseEvent>) => void;
type NodeHoverListeners = Record<'mouseover' | 'mouseout', NodeHoverListener> | {};

export interface INodeHover {
  isHovered: boolean;
  listeners: NodeHoverListeners;
}

export function useNodeHover(options: INodeHoverOptions = {}): INodeHover {
  const isDisabled = computed(() => toValue(options.isDisabled) ?? false);
  const isHovered = ref(false);

  function createListener(toHovered: boolean): NodeHoverListener {
    return () => {
      if (!isDisabled.value) isHovered.value = toHovered;
    };
  }

  whenever(isDisabled, () => isHovered.value = false);

  return reactive({
    isHovered,

    listeners: {
      mouseover: createListener(true),
      mouseout: createListener(false),
    },
  });
}
