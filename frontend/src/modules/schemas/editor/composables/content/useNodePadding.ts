import { computed, type MaybeRefOrGetter, toValue } from 'vue';
import { reactiveComputed } from '@vueuse/core';
import { getObjectKeys } from '@/helpers';

type NodePaddingSide = 'top' | 'right' | 'bottom' | 'left';
type NodePadding = Record<NodePaddingSide, number>;
type NodePaddingInputSide = NodePaddingSide | 'horizontal' | 'vertical';
export type NodePaddingInput = MaybeRefOrGetter<Partial<Record<NodePaddingInputSide, number>> | number>;

export interface INodePadding extends NodePadding {
  horizontal: number;
  vertical: number;
}

export function useNodePadding(inputRef: NodePaddingInput): INodePadding {
  const padding = computed(() => {
    const input = toValue(inputRef);

    if (typeof input === 'number') {
      return {
        top: input,
        right: input,
        bottom: input,
        left: input,
      };
    }

    return getObjectKeys(input).reduce((padding, side): NodePadding => {
      if (side === 'horizontal') {
        return { ...padding, left: input.horizontal!, right: input.horizontal! };
      }
      if (side === 'vertical') {
        return { ...padding, top: input.vertical!, bottom: input.vertical! };
      }
      return { ...padding, [side]: input[side] };
    }, {
      top: 0,
      right: 0,
      bottom: 0,
      left: 0,
    });
  });

  return reactiveComputed(() => ({
    ...padding.value,
    horizontal: padding.value.left + padding.value.right,
    vertical: padding.value.top + padding.value.bottom,
  }));
}
