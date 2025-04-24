import {
  type Component,
  computed,
  type HTMLAttributes,
  type IntrinsicElementAttributes,
} from 'vue';
import type { ComponentAs, ComponentTag, InferComponentProps } from '@/types';

export interface IBinding<P extends object> {
  is: ComponentAs;
  props: P;
}

export type ComponentBinding<C extends Component> = IBinding<Omit<HTMLAttributes, 'is'> & InferComponentProps<C>>;
export type TagBinding<T extends ComponentTag> = IBinding<IntrinsicElementAttributes[T]>;
export type Binding<C extends ComponentTag | Component> = C extends ComponentTag ? TagBinding<C> : ComponentBinding<Extract<C, Component>>;
export type AnyBinding = Binding<ComponentTag | Component>;

export function makeBinding<C extends ComponentTag | Component, const P extends Binding<C>['props'] = Binding<C>['props']>(is: C, props: () => P) {
  return computed(() => ({ is, props: props() }));
}
