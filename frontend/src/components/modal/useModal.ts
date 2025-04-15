import type { Component } from 'vue';
import type { ComponentProps } from '@/types';
import { ModalPlugin } from './ModalPlugin';

export interface IModal<C extends Component> {
  open(props?: ComponentProps<C>): void;
}

export function useModal<C extends Component>(component: C): IModal<C> {
  const plugin = ModalPlugin.inject();
  const modal = plugin.create(component);

  function open(props?: ComponentProps<C>): void {
    modal.open(props ?? null);
  }

  return { open };
}
