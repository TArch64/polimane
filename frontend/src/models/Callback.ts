import type { SafeAny } from '@/types';

export type CallbackListener<A extends SafeAny[]> = (...args: A) => void;

export class Callback<A extends SafeAny[] = []> {
  private listeners: CallbackListener<A>[] = [];

  listen(listener: CallbackListener<A>): void {
    this.listeners.push(listener);
  }

  dispatch(...args: A): void {
    for (const listener of this.listeners) {
      listener(...args);
    }
  }
}
