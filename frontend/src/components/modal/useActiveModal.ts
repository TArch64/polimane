import { type Component, inject, type InjectionKey, provide } from 'vue';
import type { Modal } from './Modal';

const PROVIDER = Symbol('ActiveModal') as InjectionKey<Modal>;

export const provideActiveModal = (modal: Modal) => provide(PROVIDER, modal);
export const useActiveModal = <R = null>() => inject(PROVIDER) as Modal<Component, R>;
