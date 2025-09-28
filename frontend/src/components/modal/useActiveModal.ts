import { type Component, inject, type InjectionKey, provide } from 'vue';
import type { ModalModel } from './ModalModel';

const PROVIDER = Symbol('ActiveModal') as InjectionKey<ModalModel>;

export const provideActiveModal = (modal: ModalModel) => provide(PROVIDER, modal);
export const useActiveModal = <R = null>() => inject(PROVIDER) as ModalModel<Component, R>;
