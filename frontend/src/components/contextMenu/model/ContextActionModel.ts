import { NodeRect } from '@/models';
import { ContextItemModel, type IContextMenuItem } from './ContextItemModel';

export interface IContextMenuEvent {
  menuRect: NodeRect;
}

export type ContextMenuOnAction = (event: IContextMenuEvent) => void | Promise<void>;

export interface IContextMenuAction extends IContextMenuItem {
  danger?: boolean;
  onAction: ContextMenuOnAction;
}

export class ContextActionModel extends ContextItemModel implements IContextMenuAction {
  readonly danger: boolean;
  readonly onAction: ContextMenuOnAction;

  constructor(def: IContextMenuAction) {
    super(def);
    this.danger = def.danger ?? false;
    this.onAction = def.onAction;
  }
}
