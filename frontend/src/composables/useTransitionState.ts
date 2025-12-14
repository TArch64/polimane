import { reactive, ref, type TransitionProps } from 'vue';
import type { PascalToKebab, SafeAny } from '@/types';

export type TransitionStateListeners = Partial<{
  [K in keyof TransitionProps as K extends `on${infer U}` ? PascalToKebab<U> : never]: Exclude<TransitionProps[K], SafeAny[]>;
}>;

export interface ITransitionState {
  isActive: boolean;
  on: () => void;
  listeners: TransitionStateListeners;
}

export function useTransitionState(): ITransitionState {
  const isActive = ref(false);
  const on = () => isActive.value = true;
  const off = () => isActive.value = false;

  const listeners: TransitionStateListeners = {
    'enter': on,
    'after-enter': off,
    'enter-cancelled': off,

    'leave': on,
    'after-leave': off,
    'leave-cancelled': off,

    'appear': on,
    'after-appear': off,
    'appear-cancelled': off,
  };

  return reactive({
    isActive,
    on,
    listeners,
  });
}
