import { reactive, type Ref } from 'vue';

export interface IConfirmButton {
  text: string;
  danger?: boolean;
}

export interface IConfirmOptions {
  id: string;
  message: string;
  anchorEl: Ref<HTMLElement | null>;
  declineButton?: string | IConfirmButton;
  acceptButton?: string | IConfirmButton;
}

interface IConfirmState {
  anchorEl: HTMLElement | null;
  isOpened: boolean;
}

export class Confirm {
  readonly id;
  readonly message;
  readonly declineButton: IConfirmButton;
  readonly acceptButton: IConfirmButton;

  private readonly state: IConfirmState;

  private promise!: Promise<boolean>;
  private resolvePromise!: (accepted: boolean) => void;

  constructor(options: IConfirmOptions) {
    this.id = options.id;
    this.message = options.message;
    this.declineButton = this.normalizeButton(options.declineButton ?? { text: 'Відмінити' });
    this.acceptButton = this.normalizeButton(options.acceptButton ?? { text: 'Підтвердити' });

    this.state = reactive({
      anchorEl: options.anchorEl,
      isOpened: false,
    });
  }

  get htmlId(): string {
    return `confirm-popover-${this.id}`;
  }

  get anchorVar(): string {
    return `--${this.htmlId}`;
  }

  get anchorEl(): HTMLElement | null {
    return this.state.anchorEl;
  }

  get isOpened(): boolean {
    return this.state.isOpened;
  }

  private normalizeButton(button: string | IConfirmButton): IConfirmButton {
    return typeof button === 'object' ? button : { text: button };
  }

  ask(): Promise<boolean> {
    this.promise = new Promise<boolean>((resolve) => {
      this.resolvePromise = resolve;
    });

    this.state.isOpened = true;
    return this.promise;
  }

  complete(accepted: boolean): void {
    this.state.isOpened = false;
    this.resolvePromise(accepted);
  }
}
