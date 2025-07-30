import { type MaybeRefOrGetter, onBeforeUnmount, watch } from 'vue';
import { toRef } from '@vueuse/core';

export function usePageClass(classNameRef: MaybeRefOrGetter<string>) {
  const className = toRef(classNameRef);
  const targetEl = document.scrollingElement!;

  watch(className, (className, oldClassName) => {
    if (className) targetEl.classList.add(className);
    if (oldClassName) targetEl.classList.remove(oldClassName);
  }, { immediate: true });

  onBeforeUnmount(() => {
    if (className.value) {
      targetEl.classList.remove(className.value);
    }
  });
}
