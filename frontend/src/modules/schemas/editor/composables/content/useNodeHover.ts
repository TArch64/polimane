import { reactive, ref } from 'vue';
import type { KonvaEventObject } from 'vue-konva';

type NodeHoverListener = (event: KonvaEventObject<MouseEvent>) => void;

export interface INodeHover {
  isHovered: boolean;
  listeners: Record<'mouseover' | 'mouseout', NodeHoverListener>;
}

export function useNodeHover(): INodeHover {
  const isHovered = ref(false);

  const onMouseOver: NodeHoverListener = () => isHovered.value = true;
  const onMouseOut: NodeHoverListener = () => isHovered.value = false;

  return reactive({
    isHovered,

    listeners: {
      mouseover: onMouseOver,
      mouseout: onMouseOut,
    },
  });
}
