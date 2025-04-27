import { reactive, ref } from 'vue';
import type { KonvaEventObject } from 'vue-konva';
import Konva from 'konva';

type NodeHoverListener = (event: KonvaEventObject<MouseEvent>) => void;

export interface INodeHover {
  isHovered: boolean;
  listeners: Record<'mouseover' | 'mouseout', NodeHoverListener>;
}

type Cursor = 'pointer' | 'default';

export interface INodeHoverOptions {
  cursor?: Cursor;
}

export function useNodeHover(options: INodeHoverOptions = {}): INodeHover {
  const isHovered = ref(false);

  function setCursor(event: KonvaEventObject<MouseEvent>, cursor: Cursor): void {
    const stage: Konva.Stage = event.target.getStage();
    stage.container().style.setProperty('cursor', cursor);
  }

  const onMouseOver: NodeHoverListener = (event) => {
    isHovered.value = true;
    if (options.cursor) setCursor(event, options.cursor);
  };

  const onMouseOut: NodeHoverListener = (event) => {
    isHovered.value = false;
    if (options.cursor) setCursor(event, 'default');
  };

  return reactive({
    isHovered,

    listeners: {
      mouseover: onMouseOver,
      mouseout: onMouseOut,
    },
  });
}
