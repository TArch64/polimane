import { shallowReactive } from 'vue';
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
  title: string;
  position: Point;
  control?: boolean;
  actions: IContextMenuAction[];
}

interface IState {
  menuRect?: NodeRect;
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
    this.actions = options.actions;
  }

  get menuRect(): NodeRect | undefined {
    return this.state.menuRect;
  }

  set menuRect(rect: NodeRect | undefined) {
    this.state.menuRect = rect;
  }

  get htmlId(): string {
    return `context-menu-popover-${this.id}`;
  }

  get anchorVar(): string {
    return `--${this.htmlId}`;
  }

  executeAction(action: IContextMenuAction): void {
    action.onAction({ menuRect: this.menuRect! });
  }
}
