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
  const on = () => isActive.value = true;
  const off = () => isActive.value = false;

  return reactive({
    isActive,

    listeners: {
      'before-enter': on,
      'after-enter': off,
      'enter-cancelled': off,

      'before-leave': on,
      'after-leave': off,
      'leave-cancelled': off,

      'before-appear': on,
      'after-appear': off,
      'appear-cancelled': off,
    },
  });
}
