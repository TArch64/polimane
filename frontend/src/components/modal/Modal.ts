import { type Component, reactive } from 'vue';
import type { ComponentProps } from '@/types';

interface IModalState<P> {
  isOpened: boolean;
  props: P | null;
}

export type ModalCloseCallback = () => Promise<void>;

export class Modal<C extends Component = Component, P = ComponentProps<C>> {
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
    this.onClose = async () => {
      await onClose?.();
      this.state.props = null;
    };

    this.state.isOpened = false;
  }
}
