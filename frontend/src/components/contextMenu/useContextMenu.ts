import { type MaybeRefOrGetter, toValue } from 'vue';
import { useEventListener } from '@vueuse/core';
import { Point } from '@/models';
import type { IContextMenuAction, MaybeContextMenuAction } from './ContextMenuModel';
import { ContextMenuPlugin } from './ContextMenuPlugin';

export function useContextMenu(el: MaybeRefOrGetter<HTMLElement>, actions: MaybeRefOrGetter<MaybeContextMenuAction[]>) {
  const plugin = ContextMenuPlugin.inject();

  useEventListener(el, 'contextmenu', (event) => {
    event.preventDefault();
    event.stopPropagation();

    plugin.show({
      position: new Point({
        x: event.clientX,
        y: event.clientY,
      }),

      actions: toValue(actions).filter((action): action is IContextMenuAction => !!action),
    });

    addEventListener('click', () => plugin.hide(), { once: true, capture: true });
  });
}
