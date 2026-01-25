import { type App, type Component, type ComponentCustomOptions, reactive } from 'vue';
import type { InferComponentProps, MaybePromise, SafeAny } from '@/types';
import { createWaiter, type IWaiter, type WaiterRelease } from '@/helpers';

interface IModalState<P> {
  isOpened: boolean;
  props: P | null;
}

export type ModalCloseCallback = () => MaybePromise<void>;

export type AnyModalModel = ModalModel<Component, SafeAny, SafeAny>;

export interface IModalOptions<C extends Component = Component> {
  id: string;
  app: App;
  component: C;
}

export class ModalModel<C extends Component = Component, R = null, P = InferComponentProps<C>> {
  private state: IModalState<P> = reactive({
    isOpened: false,
    props: null,
  });

  readonly id;
  readonly component;
  private readonly app;

  onClose?: ModalCloseCallback;
  private closeWaiter?: IWaiter;
  private closeResult?: R;
  private releaseCloseWaiter?: WaiterRelease;

  constructor(options: IModalOptions<C>) {
    this.id = options.id;
    this.component = options.component;
    this.app = options.app;
  }

  get isOpened(): boolean {
    return this.state.isOpened;
  }

  get props(): P | null {
    return this.state.props;
  }

  async open(props: P | null): Promise<R> {
    this.state.props = props;
    await this.execBeforeOpen();

    this.state.isOpened = true;

    this.closeWaiter = createWaiter();
    this.releaseCloseWaiter = this.closeWaiter.add();
    await this.closeWaiter.wait();
    return this.closeResult!;
  }

  private async execBeforeOpen(): Promise<void> {
    const beforeOpen = (this.component as ComponentCustomOptions).beforeModalOpen;
    if (!beforeOpen) return;

    await this.app.runWithContext(async () => {
      await beforeOpen(this.props ?? {});
    });
  }

  close(result: R, onClose?: ModalCloseCallback) {
    this.onClose = onClose;
    this.closeResult = result;
    this.state.isOpened = false;
  }

  async completeClose(): Promise<void> {
    await this.onClose?.();
    this.releaseCloseWaiter?.();
    this.state.props = null;
  }
}
