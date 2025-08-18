import Konva from 'konva';
import { type MaybeRefOrGetter } from 'vue';
import { useCanvasStage } from './useCanvasStage';
import { useNodeHover } from './useNodeHover';

export function useNodeCursor(nodeRef: MaybeRefOrGetter<Konva.Node>, cursor: string) {
  const stage = useCanvasStage();

  function setCursor(cursor: string) {
    stage.value?.container().style.setProperty('cursor', cursor);
  }

  useNodeHover(nodeRef, (isHovered) => {
    isHovered ? setCursor(cursor) : setCursor('default');
  });
}
