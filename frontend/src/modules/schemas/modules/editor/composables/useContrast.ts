import { computed, type MaybeRefOrGetter, reactive, toValue } from 'vue';
import { contrastWCAG21 } from 'colorjs.io/fn';

export const AA_MIN_CONTRAST = 4.5;
export const AA_HIGH_CONTRAST = 7.0;

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
  const isAA = computed(() => value.value >= AA_MIN_CONTRAST);

  return reactive({
    value: value,
    isAA: isAA,
  });
}
