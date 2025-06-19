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

  get openedModal(): Modal | null {
    return this.state.modals.find((modal) => modal.isOpened) ?? null;
  }

  create<C extends Component>(component: C): Modal<C> {
    const modal = new Modal(newId(), markRaw(component));
    this.state.modals.push(modal);
    return modal;
  }

  remove(modal: Modal): void {
    const index = this.state.modals.findIndex((m) => m.id === modal.id);
    this.state.modals.splice(index, 1);
  }
}
