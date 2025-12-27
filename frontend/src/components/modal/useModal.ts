import { type Component, onUnmounted } from 'vue';
import type { InferComponentProps } from '@/types';
import { ModalPlugin } from './ModalPlugin';

export interface IModal<C extends Component, R = null> {
  open: (props?: InferComponentProps<C>) => Promise<R>;
}

export function useModal<C extends Component, R = null>(component: C): IModal<C, R> {
  const plugin = ModalPlugin.inject();
  const modal = plugin.create<C, R>(component);

  function open(props?: InferComponentProps<C>): Promise<R> {
    return modal.open(props ?? null);
  }

  onUnmounted(() => plugin.remove(modal));

  return { open };
}
