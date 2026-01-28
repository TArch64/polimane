import { type Component, computed, type MaybeRefOrGetter, toValue } from 'vue';
import type { IModal } from '@/components/modal';
import type { InferComponentProps, MaybePromise } from '@/types';
import type { ICounter } from './counters';

export interface ILimitedActionOptions<M extends Component> {
  counter: ICounter;
  overflow?: MaybeRefOrGetter<number>;
  modal: IModal<M, boolean>;
  onAction: () => MaybePromise<void>;
}

export type LimitedAction<M extends Component> = (props: InferComponentProps<M>) => Promise<void>;

export function useLimitedAction<M extends Component>(options: ILimitedActionOptions<M>): LimitedAction<M> {
  const overflow = computed(() => toValue(options.overflow) ?? 1);

  function willOverflow(): boolean {
    return options.counter.willOverlow(overflow.value);
  }

  async function tryUpgrade(props: InferComponentProps<M>): Promise<void> {
    if (willOverflow()) {
      const isUpgraded = await options.modal.open(props);
      if (!isUpgraded) return;
    }

    if (willOverflow()) {
      return tryUpgrade(props);
    }

    await options.onAction();
  }

  return async (props) => {
    await tryUpgrade(props);
  };
}
