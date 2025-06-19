import { type MaybeRefOrGetter } from 'vue';
import Konva from 'konva';
import { useNodeListener } from './useNodeListener';
import { useNodeStage } from './useNodeStage';

export function useNodeCursor(nodeRef: MaybeRefOrGetter<Konva.Node>, cursor: string) {
  const stage = useNodeStage(nodeRef);

  function setCursor(cursor: string) {
    if (stage.value) {
      stage.value.container().style.cursor = cursor;
    }
  }

  useNodeListener(nodeRef, 'mouseover', () => setCursor(cursor));
  useNodeListener(nodeRef, 'mouseout', () => setCursor('default'));
}
