import { reactive } from 'vue';
import { NodeRect } from '@/models';

export interface IConfirmOptions {
  id: string;
  message: string;
  danger?: boolean;
  declineButton?: string;
  acceptButton?: string;
}

export interface IConfirmAskOptions {
  virtualTarget?: NodeRect;
}

interface IConfirmState {
  isOpened: boolean;
  virtualTarget?: NodeRect;
}

export class Confirm {
  readonly id;
  readonly message;
  readonly danger;
  readonly declineButton;
  readonly acceptButton;

  private readonly state: IConfirmState;

  private promise!: Promise<boolean>;
  private resolvePromise!: (accepted: boolean) => void;

  constructor(options: IConfirmOptions) {
    this.id = options.id;
    this.message = options.message;
    this.danger = options.danger ?? false;
    this.declineButton = options.declineButton ?? 'Відмінити';
    this.acceptButton = options.acceptButton ?? 'Підтвердити';

    this.state = reactive({
      isOpened: false,
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

  get virtualTarget(): NodeRect | undefined {
    return this.state.virtualTarget;
  }

  ask(options: IConfirmAskOptions = {}): Promise<boolean> {
    this.promise = new Promise<boolean>((resolve) => {
      this.resolvePromise = resolve;
    });

    this.state.virtualTarget = options.virtualTarget;
    this.state.isOpened = true;
    return this.promise;
  }

  complete(accepted: boolean): void {
    this.state.isOpened = false;
    this.resolvePromise(accepted);
  }
}
