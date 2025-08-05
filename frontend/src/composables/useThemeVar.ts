import { computed, type ComputedRef, type MaybeRefOrGetter, shallowRef, toValue } from 'vue';

type KnownColor
  = 'white'
    | 'black'
    | 'danger'
    | 'primary'
    | 'divider'
    | 'hover-divider'
    | 'background-1'
    | 'background-2'
    | 'background-3'
    | 'text-1'
    | 'text-2'
    | 'text-3';

type KnownFont = 'xs' | 'sm' | 'md';
type KnownRounded = 'sm' | 'md' | 'full';

type PxVars = `--font-${KnownFont}` | `--rounded-${KnownRounded}`;
type ColorVars = `--color-${KnownColor}`;

const cache = shallowRef<CSSStyleDeclaration>(null!);

export function useThemeVar(nameRef: MaybeRefOrGetter<PxVars>): ComputedRef<number>;
export function useThemeVar(nameRef: MaybeRefOrGetter<ColorVars | string>): ComputedRef<string>;
export function useThemeVar(nameRef: MaybeRefOrGetter<string>): ComputedRef<string | number> {
  cache.value ??= getComputedStyle(document.documentElement);

  return computed(() => {
    const name = toValue(nameRef);
    const value = cache.value.getPropertyValue(name);

    if (name.endsWith('px')) {
      return parseFloat(value.trim());
    }

    return value.trim();
  });
}
