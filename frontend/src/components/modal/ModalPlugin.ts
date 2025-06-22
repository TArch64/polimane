import {
  type Component,
  type FunctionPlugin,
  inject,
  type InjectionKey,
  markRaw,
  reactive,
} from 'vue';
import { newId } from '@/helpers';
import { type AnyModal, Modal } from './Modal';

const PROVIDER = Symbol('ModalPlugin') as InjectionKey<ModalPlugin>;

interface IModalPluginState {
  modals: AnyModal[];
}

export class ModalPlugin {
  static install: FunctionPlugin = (app) => {
    app.provide(PROVIDER, new ModalPlugin());
  };

  static inject(): ModalPlugin {
    return inject(PROVIDER)!;
  }

  private state: IModalPluginState = reactive({
    modals: [],
  });

  get openedModal(): AnyModal | null {
    return this.state.modals.find((modal) => modal.isOpened) ?? null;
  }

  create<C extends Component, R = null>(component: C): Modal<C, R> {
    const modal = new Modal<C, R>(newId(), markRaw(component));
    this.state.modals.push(modal);
    return modal;
  }

  remove(modal: AnyModal): void {
    const index = this.state.modals.findIndex((m) => m.id === modal.id);
    this.state.modals.splice(index, 1);
  }
}
