import { computed, type MaybeRefOrGetter, reactive, ref, toValue, watch } from 'vue';
import type { KonvaEventObject } from 'vue-konva';

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

  const onMouseOver: NodeHoverListener = () => isHovered.value = true;
  const onMouseOut: NodeHoverListener = () => isHovered.value = false;

  const listeners = computed((): NodeHoverListeners => {
    return isDisabled.value ? {} : { mouseover: onMouseOver, mouseout: onMouseOut };
  });

  watch(isDisabled, (isDisabled) => {
    if (isDisabled) isHovered.value = false;
  });

  return reactive({
    isHovered,
    listeners,
  });
}
