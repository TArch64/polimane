import { shallowReactive } from 'vue';
import { Callback, NodeRect, Point } from '@/models';
import type { MaybeValue } from '@/types';
import { ContextActionModel, type IContextMenuAction } from './ContextActionModel';
import { ContextGroupModel, type IContextMenuGroup } from './ContextGroupModel';

export function isContextMenuAction(item: ContextMenuItem): item is IContextMenuAction {
  return (item as IContextMenuAction).onAction !== undefined;
}

export function isContextMenuGroup(item: ContextMenuItem): item is IContextMenuGroup {
  return (item as IContextMenuGroup).actions !== undefined;
}

export type ContextMenuItem = IContextMenuAction | IContextMenuGroup;
export type ContextMenuItemModel = ContextActionModel | ContextGroupModel;
export type MaybeContextMenuAction = MaybeValue<ContextMenuItem>;

export interface IContextMenuOptions {
  id: string;
  title: string;
  position: Point;
  control?: boolean;
  actions: ContextMenuItem[];
}

interface IState {
  menuRect?: NodeRect;
  openedGroup?: ContextGroupModel;
}

export class ContextMenuModel {
  readonly id;
  readonly title;
  readonly position;
  readonly control;
  readonly actions;
  readonly onHide = new Callback();

  private readonly state = shallowReactive<IState>({});

  constructor(options: IContextMenuOptions) {
    this.id = options.id;
    this.title = options.title;
    this.position = options.position;
    this.control = options.control ?? true;

    this.actions = options.actions.map((item): ContextMenuItemModel => {
      return isContextMenuAction(item)
        ? new ContextActionModel(item)
        : new ContextGroupModel(item);
    });
  }

  get menuRect(): NodeRect | undefined {
    return this.state.menuRect;
  }

  setMenuRect(rect: NodeRect | undefined) {
    this.state.menuRect = rect;
  }

  get htmlId(): string {
    return `context-menu-popover-${this.id}`;
  }

  get anchorVar(): string {
    return `--${this.htmlId}`;
  }

  get openedGroup(): ContextGroupModel | undefined {
    return this.state.openedGroup;
  }

  executeAction(action: ContextActionModel): void {
    action.onAction({ menuRect: this.menuRect! });
  }

  openGroup(group: ContextGroupModel): void {
    this.state.openedGroup = group;
  }

  closeGroup(): void {
    this.state.openedGroup = undefined;
  }
}
