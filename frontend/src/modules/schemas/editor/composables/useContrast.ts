import { computed, type MaybeRefOrGetter, toValue } from 'vue';
import { contrastWCAG21 } from 'colorjs.io/fn';

export function useContrast(
  backgroundRef: MaybeRefOrGetter<string>,
  foregroundRef: MaybeRefOrGetter<string>,
) {
  const background = computed(() => toValue(backgroundRef));
  const foreground = computed(() => toValue(foregroundRef));
  return computed(() => contrastWCAG21(background.value, foreground.value));
}
