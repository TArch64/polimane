import type { MaybeRef } from 'vue';

export type WrapMaybeRef<O extends object, P extends keyof O> = {
  [K in keyof O]: K extends P ? O[K] | MaybeRef<O[K]> : O[K];
};
