import type { IconComponent } from '@/components/icon';
import { newId } from '@/helpers';

export interface IContextMenuItem {
  title: string;
  icon: IconComponent;
  disabled?: boolean;
}

export abstract class ContextItemModel implements IContextMenuItem {
  readonly id;
  readonly title;
  readonly icon;
  readonly disabled;

  protected constructor(def: IContextMenuItem) {
    this.id = newId();
    this.title = def.title;
    this.icon = def.icon;
    this.disabled = def.disabled ?? false;
  }
}
