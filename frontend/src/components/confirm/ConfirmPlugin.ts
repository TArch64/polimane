import { type FunctionPlugin, inject, type InjectionKey, reactive } from 'vue';
import { newId } from '@/helpers';
import { ConfirmModel, type IConfirmOptions } from './ConfirmModel';

const PROVIDER = Symbol('ConfirmPlugin') as InjectionKey<ConfirmPlugin>;

interface IConfirmPluginState {
  confirms: ConfirmModel[];
}

export type ConfirmCreateInternalOptions = Omit<IConfirmOptions, 'id'>;

export class ConfirmPlugin {
  static install: FunctionPlugin = (app) => {
    app.provide(PROVIDER, new ConfirmPlugin());
  };

  static inject(): ConfirmPlugin {
    return inject(PROVIDER)!;
  }

  private readonly state: IConfirmPluginState = reactive({
    confirms: [],
  });

  get openedConfirm(): ConfirmModel | undefined {
    return this.state.confirms.slice().reverse().find((confirm) => confirm.isOpened);
  }

  create(options: ConfirmCreateInternalOptions) {
    const confirm = new ConfirmModel({
      ...options,
      id: newId(),
    });

    this.state.confirms.push(confirm);
    return confirm;
  }

  remove(confirm: ConfirmModel): void {
    const index = this.state.confirms.findIndex((c) => c.id === confirm.id);
    this.state.confirms.splice(index, 1);
    confirm.markAsRemoved();
  }
}
