import { type MaybeRefOrGetter, toValue } from 'vue';
import { useEventListener } from '@vueuse/core';
import { Point } from '@/models';
import type { IContextMenuAction, MaybeContextMenuAction } from './ContextMenuModel';
import { ContextMenuPlugin } from './ContextMenuPlugin';

export interface IContextMenuViewOptions {
  el: MaybeRefOrGetter<HTMLElement>;
  title: MaybeRefOrGetter<string>;
  actions: MaybeRefOrGetter<MaybeContextMenuAction[]>;
  control?: boolean;
}

export function useContextMenu(options: IContextMenuViewOptions) {
  const plugin = ContextMenuPlugin.inject();

  useEventListener(options.el, 'contextmenu', (event) => {
    event.preventDefault();
    event.stopPropagation();

    plugin.show({
      position: new Point({
        x: event.clientX,
        y: event.clientY,
      }),

      title: toValue(options.title),
      control: options.control,
      actions: toValue(options.actions).filter((action): action is IContextMenuAction => !!action),
    });

    addEventListener('click', () => plugin.hide(), { once: true, capture: true });
  });
}
