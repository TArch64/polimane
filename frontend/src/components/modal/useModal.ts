import { type Component, onUnmounted } from 'vue';
import { createEventHook } from '@vueuse/core';
import type { EventHookOn } from '@vueuse/shared';
import type { InferComponentProps } from '@/types';
import { ModalPlugin } from './ModalPlugin';

export interface IModal<C extends Component, R = null> {
  open: (props?: InferComponentProps<C>) => Promise<R>;
  onResult: EventHookOn<[R]>;
}

export function useModal<C extends Component, R = null>(component: C): IModal<C, R> {
  const plugin = ModalPlugin.inject();
  const modal = plugin.create<C, R>(component);
  const resultHook = createEventHook<[R]>();

  async function open(props?: InferComponentProps<C>): Promise<R> {
    const result = await modal.open(props ?? null);
    if (result) await resultHook.trigger(result);
    return result;
  }

  onUnmounted(() => plugin.remove(modal));

  return { open, onResult: resultHook.on };
}
