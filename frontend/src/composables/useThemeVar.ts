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

function updateCache() {
  cache.value = getComputedStyle(document.documentElement);
}

export function useThemeVar(nameRef: MaybeRefOrGetter<PxVars>): ComputedRef<number>;
export function useThemeVar(nameRef: MaybeRefOrGetter<ColorVars | string>): ComputedRef<string>;
export function useThemeVar(nameRef: MaybeRefOrGetter<string>): ComputedRef<string | number> {
  if (!cache.value) updateCache();

  return computed(() => {
    const name = toValue(nameRef);
    const value = cache.value.getPropertyValue(name).trim();

    if (value.endsWith('px')) {
      return parseFloat(value);
    }

    return value || name;
  });
}

import.meta.hot?.on('vite:afterUpdate', (payload) => {
  const needUpdate = payload.updates.some((update) => {
    return update.type === 'js-update' && update.path.includes('style/main.css');
  });

  if (needUpdate) updateCache();
});
