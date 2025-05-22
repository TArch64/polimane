import type { IconComponent } from '@/components/icon';
import { NodeRect, Point } from '@/models';

export interface IContextMenuEvent {
  menuRect: NodeRect;
}

export type ContextMenuOnAction = (event: IContextMenuEvent) => void | Promise<void>;

export interface IContextMenuAction {
  title: string;
  icon: IconComponent;
  danger?: boolean;
  onAction: ContextMenuOnAction;
}

export type MaybeContextMenuAction = IContextMenuAction | null | undefined | false;

export interface IContextMenuOptions {
  id: string;
  position: Point;
  actions: IContextMenuAction[];
}

export class ContextMenu {
  readonly id;
  readonly position;
  readonly actions;
  private menuRect?: NodeRect;

  constructor(options: IContextMenuOptions) {
    this.id = options.id;
    this.position = options.position;
    this.actions = options.actions;
  }

  setMenuRect(rect: NodeRect): void {
    this.menuRect = rect;
  }

  executeAction(action: IContextMenuAction): void {
    action.onAction({ menuRect: this.menuRect! });
  }
}
