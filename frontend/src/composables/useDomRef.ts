import { computed, type Ref, shallowRef } from 'vue';
import { type MaybeComputedElementRef, unrefElement } from '@vueuse/core';

export function useDomRef<E extends HTMLElement | null>(): Ref<E, MaybeComputedElementRef | null> {
  const domRef = shallowRef<E>(null!);

  return computed<E, MaybeComputedElementRef | null>({
    get: () => domRef.value,
    set: (value) => domRef.value = unrefElement(value),
  });
}
