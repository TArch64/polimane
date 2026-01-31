import { type MaybeRefOrGetter, reactive } from 'vue';
import { NodeRect } from '@/models';
import type { ComponentVariant } from '@/types';

export interface IConfirmOptions {
  id: string;
  message: MaybeRefOrGetter<string>;
  topEl: MaybeRefOrGetter<HTMLElement>;
  danger?: boolean;
  variant?: ComponentVariant;
  declineButton?: string;
  acceptButton?: string;
  additionalCondition?: string;
}

export interface IConfirmAskOptions {
  virtualTarget?: NodeRect;
}

interface IConfirmState {
  isOpened: boolean;
  isRemoved: boolean;
  virtualTarget?: NodeRect;
}

export interface IConfirmResult {
  isAccepted: boolean;
  isSecondaryAccepted?: boolean;
}

export class ConfirmModel {
  readonly id;
  readonly message;
  readonly topEl;
  readonly danger;
  readonly declineButton;
  readonly acceptButton;
  readonly additionalCondition;

  private readonly state: IConfirmState;

  private promise!: Promise<IConfirmResult>;
  private resolvePromise!: (result: IConfirmResult) => void;

  constructor(options: IConfirmOptions) {
    this.id = options.id;
    this.message = options.message;
    this.topEl = options.topEl;
    this.danger = options.danger ?? false;
    this.declineButton = options.declineButton ?? 'Відмінити';
    this.acceptButton = options.acceptButton ?? 'Підтвердити';
    this.additionalCondition = options.additionalCondition ?? '';

    this.state = reactive({
      isOpened: false,
      isRemoved: false,
    });
  }

  get htmlId(): string {
    return `confirm-popover-${this.id}`;
  }

  get anchorVar(): string {
    return `--${this.htmlId}`;
  }

  get isOpened(): boolean {
    return this.state.isOpened;
  }

  get isRemoved(): boolean {
    return this.state.isRemoved;
  }

  get virtualTarget(): NodeRect | undefined {
    return this.state.virtualTarget;
  }

  ask(options: IConfirmAskOptions = {}): Promise<IConfirmResult> {
    const { promise, resolve } = Promise.withResolvers<IConfirmResult>();
    this.promise = promise;
    this.resolvePromise = resolve;

    this.state.virtualTarget = options.virtualTarget;
    this.state.isOpened = true;
    return this.promise;
  }

  complete(result: IConfirmResult): void {
    this.state.isOpened = false;
    this.resolvePromise(result);
  }

  markAsRemoved(): void {
    this.state.isRemoved = true;
  }
}
