import { markRaw } from 'vue';
import type { SafeAny } from '@/types';

export type CallbackListener<A extends SafeAny[]> = (...args: A) => void;

export class Callback<A extends SafeAny[] = []> {
  static create<A extends SafeAny[] = []>(): Callback<A> {
    return markRaw(new Callback<A>());
  }

  #listeners: CallbackListener<A>[] = [];

  listen(listener: CallbackListener<A>): () => void {
    this.#listeners.push(listener);
    return () => this.#listeners = this.#listeners.filter((l) => l !== listener);
  }

  dispatch(...args: A): void {
    for (const listener of this.#listeners) {
      listener(...args);
    }
  }
}
