import type { SafeAny } from '@/types';

export type AnimatedFrameTick = (...args: SafeAny[]) => void;

export function createAnimatedFrame<T extends AnimatedFrameTick>(tick: T): T {
  let frameId: number | null = null;

  return ((...args: SafeAny[]) => {
    if (!frameId) {
      frameId = requestAnimationFrame(() => {
        frameId = null;
        tick(...args);
      });
    }
  }) as T;
}
