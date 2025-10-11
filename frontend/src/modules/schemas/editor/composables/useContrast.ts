import { computed, type MaybeRefOrGetter, reactive, toValue } from 'vue';
import { contrastWCAG21 } from 'colorjs.io/fn';

export const MIN_CONTRAST_AA = 4.5;

export interface IContrast {
  value: number;
  isAA: boolean;
}

export function useContrast(
  backgroundRef: MaybeRefOrGetter<string>,
  foregroundRef: MaybeRefOrGetter<string>,
): IContrast {
  const background = computed(() => toValue(backgroundRef));
  const foreground = computed(() => toValue(foregroundRef));
  const value = computed(() => contrastWCAG21(background.value, foreground.value));
  const isAA = computed(() => value.value >= MIN_CONTRAST_AA);

  return reactive({
    value: value,
    isAA: isAA,
  });
}
