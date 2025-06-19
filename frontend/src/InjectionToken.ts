import type { InjectionKey, Ref } from 'vue';

export const TOKEN_SCROLLER: InjectionKey<Ref<HTMLElement>> = Symbol('scroller');
export const TOKEN_TOP_EL: InjectionKey<Ref<HTMLElement>> = Symbol('topRef');
