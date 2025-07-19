import { type FunctionPlugin, inject, type InjectionKey, shallowRef } from 'vue';
import { newId } from '@/helpers';
import { ContextMenuModel, type IContextMenuOptions } from './model';

const PROVIDER = Symbol('ContextMenuPlugin') as InjectionKey<ContextMenuPlugin>;

export class ContextMenuPlugin {
  static install: FunctionPlugin = (app) => {
    app.provide(PROVIDER, new ContextMenuPlugin());
  };

  static inject(): ContextMenuPlugin {
    return inject(PROVIDER)!;
  }

  activeMenu = shallowRef<ContextMenuModel | null>(null);

  show(options: Omit<IContextMenuOptions, 'id'>) {
    this.activeMenu.value = new ContextMenuModel({
      ...options,
      id: newId(),
    });
  }

  hide() {
    this.activeMenu.value = null;
  }
}
