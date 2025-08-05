import { type Component, reactive } from 'vue';
import type { InferComponentProps, MaybePromise, SafeAny } from '@/types';
import { createWaiter, type IWaiter, type WaiterRelease } from '@/helpers';

interface IModalState<P> {
  isOpened: boolean;
  props: P | null;
}

export type ModalCloseCallback = () => MaybePromise<void>;

export type AnyModal = Modal<Component, SafeAny, SafeAny>;

export class Modal<C extends Component = Component, R = null, P = InferComponentProps<C>> {
  private state: IModalState<P> = reactive({
    isOpened: false,
    props: null,
  });

  onClose?: ModalCloseCallback;
  private closeWaiter?: IWaiter;
  private closeResult?: R;
  private releaseCloseWaiter?: WaiterRelease;

  constructor(
    readonly id: string,
    readonly component: C,
  ) {
  }

  get isOpened(): boolean {
    return this.state.isOpened;
  }

  get props(): P | null {
    return this.state.props;
  }

  async open(props: P | null): Promise<R> {
    this.state.props = props;
    this.state.isOpened = true;

    this.closeWaiter = createWaiter();
    this.releaseCloseWaiter = this.closeWaiter.add();
    await this.closeWaiter.wait();
    return this.closeResult!;
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
