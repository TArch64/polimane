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

  show(options: Omit<IContextMenuOptions, 'id'>): ContextMenuModel {
    const menu = new ContextMenuModel({
      ...options,
      id: newId(),
    });

    this.activeMenu.value?.onHide.dispatch();
    this.activeMenu.value = menu;
    return menu;
  }

  hide() {
    this.activeMenu.value?.onHide.dispatch();
    this.activeMenu.value = null;
  }
}
