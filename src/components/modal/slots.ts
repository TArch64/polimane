import type { Slot } from 'vue';

export type ModalActivatorSlot = Slot<{
  open: () => void;
}>;

export type ModalContentSlot = Slot<{
  close: () => void;
}>;
