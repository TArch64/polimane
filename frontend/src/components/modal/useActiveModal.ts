import { inject, type InjectionKey, provide } from 'vue';
import type { Modal } from './Modal';

const PROVIDER = Symbol('ActiveModal') as InjectionKey<Modal>;

export const provideActiveModal = (modal: Modal) => provide(PROVIDER, modal);
export const useActiveModal = () => inject(PROVIDER)!;
