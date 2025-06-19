import { type Component, reactive } from 'vue';
import type { InferComponentProps } from '@/types';

interface IModalState<P> {
  isOpened: boolean;
  props: P | null;
}

export type ModalCloseCallback = () => Promise<void>;

export class Modal<C extends Component = Component, P = InferComponentProps<C>> {
  private state: IModalState<P> = reactive({
    isOpened: false,
    props: null,
  });

  onClose?: ModalCloseCallback;

  constructor(
    readonly id: string,
    readonly component: Component,
  ) {
  }

  get isOpened(): boolean {
    return this.state.isOpened;
  }

  get props(): P | null {
    return this.state.props;
  }

  open(props: P | null) {
    this.state.props = props;
    this.state.isOpened = true;
  }

  close(onClose?: ModalCloseCallback) {
    this.onClose = onClose;
    this.state.isOpened = false;
  }

  async completeClose(): Promise<void> {
    await this.onClose?.();
    this.state.props = null;
  }
}
