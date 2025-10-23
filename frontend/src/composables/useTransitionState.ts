import { reactive, ref, type TransitionProps } from 'vue';
import type { PascalToKebab } from '@/types';

export type TransitionStateListeners = Partial<{
  [K in keyof TransitionProps as K extends `on${infer U}` ? PascalToKebab<U> : never]: TransitionProps[K];
}>;

export interface ITransitionState {
  isActive: boolean;
  listeners: TransitionStateListeners;
}

export function useTransitionState(): ITransitionState {
  const isActive = ref(false);

  return reactive({
    isActive,

    listeners: {
      'before-enter': () => isActive.value = true,
      'after-enter': () => isActive.value = false,
      'before-leave': () => isActive.value = true,
      'after-leave': () => isActive.value = false,
    },
  });
}
