import { nextTick, reactive, ref, type Ref } from 'vue';
import { waitClickComplete } from '@/helpers';
import { type ITransitionState, useTransitionState } from './useTransitionState';

export interface IPopoverState {
  transition: ITransitionState;
  isOpened: boolean;
  open: () => Promise<void>;
  close: () => void;
}

export function usePopoverState(dialogRef: Ref<HTMLElement | null>): IPopoverState {
  let closeController: AbortController | null = null;
  const transition = useTransitionState();
  const isOpened = ref(false);

  function close(): void {
    closeController?.abort();
    transition.on();
    isOpened.value = false;
  }

  function closeEvent(event: Event): void {
    if (dialogRef.value?.contains(event.target as Node)) {
      return;
    }

    close();
  }

  async function open() {
    if (isOpened.value) {
      return;
    }

    isOpened.value = true;
    await nextTick();

    dialogRef.value!.showPopover();
    await waitClickComplete();

    closeController = new AbortController();

    window.addEventListener('mousedown', closeEvent, {
      signal: closeController.signal,
      capture: true,
    });

    window.addEventListener('contextmenu', closeEvent, {
      signal: closeController.signal,
      capture: true,
    });
  }

  return reactive({
    transition,
    isOpened,
    open,
    close,
  });
}
