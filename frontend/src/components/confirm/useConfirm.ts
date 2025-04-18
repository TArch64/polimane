import { onUnmounted, type VNodeRef } from 'vue';
import { useDomRef } from '@/composables';
import { type ConfirmCreateOptions, ConfirmPlugin } from './ConfirmPlugin';

export type ConfirmOptions = Omit<ConfirmCreateOptions, 'anchorEl'>;

export interface IConfirm {
  ask: () => Promise<boolean>;
  anchorRef: VNodeRef;
}

export function useConfirm(options: ConfirmOptions): IConfirm {
  const plugin = ConfirmPlugin.inject();
  const anchorEl = useDomRef<HTMLElement>();

  const confirm = plugin.create({
    ...options,
    anchorEl: anchorEl.ref,
  });

  const ask = () => confirm.ask();

  onUnmounted(() => plugin.remove(confirm));

  return { ask, anchorRef: anchorEl.templateRef };
}
