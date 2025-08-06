import { type MaybeRefOrGetter } from 'vue';
import Konva from 'konva';
import { useNodeStage } from './useNodeStage';
import { useNodeHovering } from './useNodeHover';

export function useNodeCursor(nodeRef: MaybeRefOrGetter<Konva.Node>, cursor: string) {
  const stage = useNodeStage(nodeRef);

  function setCursor(cursor: string) {
    if (stage.value) {
      stage.value.container().style.cursor = cursor;
    }
  }

  useNodeHovering(nodeRef, (isHovering) => {
    isHovering ? setCursor(cursor) : setCursor('default');
  });
}
