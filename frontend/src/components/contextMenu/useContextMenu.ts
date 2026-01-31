import { computed, type MaybeRefOrGetter, ref, type Ref, toValue } from 'vue';
import { useEventListener } from '@vueuse/core';
import { Point } from '@/models';
import type { ComponentVariant } from '@/types';
import type { ContextMenuItem, MaybeContextMenuAction } from './model';
import { ContextMenuPlugin } from './ContextMenuPlugin';

export interface IContextMenuViewOptions {
  el: MaybeRefOrGetter<HTMLElement | null>;
  title: MaybeRefOrGetter<string>;
  actions: MaybeRefOrGetter<MaybeContextMenuAction[]>;
  variant?: ComponentVariant;
  isActive?: Ref<boolean>;
}

export function useContextMenu(options: IContextMenuViewOptions) {
  const plugin = ContextMenuPlugin.inject();
  const isActive = options.isActive ?? ref(false);

  const availableActions = computed(() => {
    return toValue(options.actions)
      .filter((action): action is ContextMenuItem => !!action);
  });

  useEventListener(options.el, 'contextmenu', (event) => {
    event.preventDefault();
    event.stopPropagation();

    if (!availableActions.value.length) {
      return;
    }

    isActive.value = true;

    const model = plugin.show({
      position: new Point(event.clientX, event.clientY),
      title: toValue(options.title),
      variant: options.variant,
      actions: availableActions.value,
    });

    model.onHide.listen(() => isActive.value = false);
  });
}
