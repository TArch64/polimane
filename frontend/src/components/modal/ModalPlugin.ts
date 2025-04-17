import {
  type Component,
  type FunctionPlugin,
  inject,
  type InjectionKey,
  markRaw,
  reactive,
} from 'vue';
import { newId } from '@/helpers';
import { Modal } from './Modal';

const PROVIDER = Symbol('ModalPlugin') as InjectionKey<ModalPlugin>;

interface IModalPluginState {
  modals: Modal[];
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

  get openedModal(): Modal | undefined {
    return this.state.modals.slice().reverse().find((modal) => modal.isOpened);
  }

  create(component: Component): Modal {
    const modal = new Modal(newId(), markRaw(component));
    this.state.modals.push(modal);
    return modal;
  }
}
