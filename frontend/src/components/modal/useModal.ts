import { type Component, onUnmounted } from 'vue';
import type { InferComponentProps } from '@/types';
import { ModalPlugin } from './ModalPlugin';

export interface IModal<C extends Component> {
  open(props?: InferComponentProps<C>): void;
}

export function useModal<C extends Component>(component: C): IModal<C> {
  const plugin = ModalPlugin.inject();
  const modal = plugin.create<C>(component);

  function open(props?: InferComponentProps<C>): void {
    modal.open(props ?? null);
  }

  onUnmounted(() => plugin.remove(modal));

  return { open };
}
