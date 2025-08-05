import { computed, type ComputedRef, type MaybeRefOrGetter, shallowRef, toValue } from 'vue';

const cache = shallowRef<CSSStyleDeclaration>(null!);

export function useThemeVar(nameRef: MaybeRefOrGetter<string>): ComputedRef<string> {
  cache.value ??= getComputedStyle(document.documentElement);

  return computed(() => {
    const name = toValue(nameRef);
    return cache.value.getPropertyValue(name);
  });
}
