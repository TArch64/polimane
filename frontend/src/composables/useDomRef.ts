import { type ShallowRef, shallowRef, type VNodeRef } from 'vue';
import { type MaybeElement, unrefElement } from '@vueuse/core';

export interface IDomRef<E extends HTMLElement> {
  ref: ShallowRef<E>;
  templateRef: VNodeRef;
}

export function useDomRef<E extends HTMLElement>(): IDomRef<E> {
  const domRef = shallowRef<E>(null!);

  const templateRef: VNodeRef = (ref) => {
    domRef.value = unrefElement(ref as MaybeElement);
  };

  return { ref: domRef as ShallowRef<E>, templateRef };
}
