import { shallowReactive } from 'vue';
import type { IconComponent } from '@/components/icon';
import { NodeRect, Point } from '@/models';
import { newId } from '@/helpers';

export interface IContextMenuEvent {
  menuRect: NodeRect;
}

export type ContextMenuOnAction = (event: IContextMenuEvent) => void | Promise<void>;

interface IContextMenuItem {
  id: string;
  title: string;
  icon: IconComponent;
}

export interface IContextMenuAction extends IContextMenuItem {
  danger?: boolean;
  onAction: ContextMenuOnAction;
}

export function isContextMenuAction(item: ContextMenuItem): item is IContextMenuAction {
  return (item as IContextMenuAction).onAction !== undefined;
}

export interface IContextMenuGroup extends IContextMenuItem {
  actions: IContextMenuAction[];
}

export function isContextMenuGroup(item: ContextMenuItem): item is IContextMenuGroup {
  return (item as IContextMenuGroup).actions !== undefined;
}

export type ContextMenuItem = IContextMenuAction | IContextMenuGroup;
export type ContextMenuItemDefinition = Omit<ContextMenuItem, 'id'>;
export type MaybeContextMenuAction = ContextMenuItemDefinition | null | undefined | false;

export interface IContextMenuOptions {
  id: string;
  title: string;
  position: Point;
  control?: boolean;
  actions: ContextMenuItemDefinition[];
}

interface IState {
  menuRect?: NodeRect;
  openedGroup?: IContextMenuGroup;
}

export class ContextMenuModel {
  readonly id;
  readonly title;
  readonly position;
  readonly control;
  readonly actions;

  private readonly state = shallowReactive<IState>({});

  constructor(options: IContextMenuOptions) {
    this.id = options.id;
    this.title = options.title;
    this.position = options.position;
    this.control = options.control ?? true;

    this.actions = options.actions.map((def) => ({
      ...def,
      id: newId(),
    }) as ContextMenuItem);
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

  get openedGroup(): IContextMenuGroup | undefined {
    return this.state.openedGroup;
  }

  executeAction(action: IContextMenuAction): void {
    action.onAction({ menuRect: this.menuRect! });
  }

  openGroup(group: IContextMenuGroup): void {
    this.state.openedGroup = group;
  }

  closeGroup(): void {
    this.state.openedGroup = undefined;
  }
}
