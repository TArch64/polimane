import { ContextActionModel, type IContextMenuAction } from './ContextActionModel';
import { ContextItemModel, type IContextMenuItem } from './ContextItemModel';

export interface IContextMenuGroup extends IContextMenuItem {
  actions: IContextMenuAction[];
}

export class ContextGroupModel extends ContextItemModel implements IContextMenuGroup {
  readonly actions: ContextActionModel[];

  constructor(def: IContextMenuGroup) {
    const disabled = def.disabled ?? def.actions.every((action) => action.disabled);
    super({ ...def, disabled });

    this.actions = def.actions.map((action) => new ContextActionModel(action));
  }
}
