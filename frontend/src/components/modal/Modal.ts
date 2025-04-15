import { type Component, reactive } from 'vue';
import type { ComponentProps } from '@/types';

interface IModalState<P> {
  isOpened: boolean;
  props: P | null;
}

export class Modal<C extends Component = Component, P = ComponentProps<C>> {
  private state: IModalState<P> = reactive({
    isOpened: false,
    props: null,
  });

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

  close() {
    this.state.isOpened = false;
    this.state.props = null;
  }
}
