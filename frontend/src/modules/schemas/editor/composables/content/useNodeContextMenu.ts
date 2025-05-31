import {
  computed,
  type MaybeRefOrGetter,
  nextTick,
  onBeforeUnmount,
  onMounted,
  toValue,
} from 'vue';
import Konva from 'konva';
import type { KonvaEventObject } from 'konva/lib/Node';
import {
  ContextMenuPlugin,
  type IContextMenuAction,
  type MaybeContextMenuAction,
} from '@/components/contextMenu';
import { Point } from '@/models';

export function useNodeContextMenu(
  nodeRef: MaybeRefOrGetter<Konva.Node>,
  actions: MaybeRefOrGetter<MaybeContextMenuAction[]>,
): void {
  const plugin = ContextMenuPlugin.inject();

  const node = computed(() => toValue(nodeRef));
  const stage = computed(() => node.value?.getStage());

  function isClosestCurrentNode(target: Konva.Stage | Konva.Node): boolean {
    if (target._id === node.value._id) {
      return true;
    }

    return target.parent ? isClosestCurrentNode(target.parent) : false;
  }

  function onContextMenu(event: KonvaEventObject<MouseEvent>): void {
    if (event.cancelBubble) {
      return;
    }

    if (!isClosestCurrentNode(event.target)) {
      return;
    }

    event.cancelBubble = true;

    plugin.show({
      position: new Point({
        x: event.evt.clientX,
        y: event.evt.clientY,
      }),

      actions: toValue(actions).filter((action): action is IContextMenuAction => !!action),
    });

    addEventListener('click', () => plugin.hide(), { once: true, capture: true });
  }

  onMounted(async () => {
    await nextTick();
    stage.value?.on('contextmenu', onContextMenu);
  });

  onBeforeUnmount(() => {
    stage.value?.off('contextmenu', onContextMenu);
  });
}
