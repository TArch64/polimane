import { type FunctionPlugin, inject, type InjectionKey, reactive } from 'vue';
import { newId } from '@/helpers';
import { Confirm, type IConfirmOptions } from './Confirm';

const PROVIDER = Symbol('ConfirmPlugin') as InjectionKey<ConfirmPlugin>;

interface IConfirmPluginState {
  confirms: Confirm[];
}

export type ConfirmCreateOptions = Omit<IConfirmOptions, 'id'>;

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

  get openedConfirm(): Confirm | undefined {
    return this.state.confirms.slice().reverse().find((confirm) => confirm.isOpened);
  }

  create(options: ConfirmCreateOptions) {
    const confirm = new Confirm({
      ...options,
      id: newId(),
    });

    this.state.confirms.push(confirm);
    return confirm;
  }
}
