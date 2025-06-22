import { type Component, onUnmounted } from 'vue';
import { createEventHook } from '@vueuse/core';
import type { EventHookOn } from '@vueuse/shared';
import type { InferComponentProps } from '@/types';
import { ModalPlugin } from './ModalPlugin';

export interface IModal<C extends Component, R = null> {
  open: (props?: InferComponentProps<C>) => Promise<R>;
  onClose: EventHookOn<[R]>;
}

export function useModal<C extends Component, R = null>(component: C): IModal<C, R> {
  const plugin = ModalPlugin.inject();
  const modal = plugin.create<C, R>(component);
  const closeHook = createEventHook<[R]>();

  async function open(props?: InferComponentProps<C>): Promise<R> {
    const result = await modal.open(props ?? null);
    await closeHook.trigger(result);
    return result;
  }

  onUnmounted(() => plugin.remove(modal));

  return { open, onClose: closeHook.on };
}
