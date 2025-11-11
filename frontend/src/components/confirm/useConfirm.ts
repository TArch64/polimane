import { onBeforeUnmount } from 'vue';
import { useTopElement } from '@/composables';
import { type ConfirmCreateInternalOptions, ConfirmPlugin } from './ConfirmPlugin';
import type { IConfirmAskOptions } from './Confirm';

export interface IConfirm {
  ask: (options?: IConfirmAskOptions) => Promise<boolean>;
  anchorStyle: { anchorName: string };
}

export type ConfirmCreateOptions = Omit<ConfirmCreateInternalOptions, 'topEl'>;

export function useConfirm(options: ConfirmCreateOptions): IConfirm {
  const plugin = ConfirmPlugin.inject();

  const confirm = plugin.create({
    ...options,
    topEl: useTopElement(),
  });

  const ask = (options: IConfirmAskOptions = {}) => confirm.ask(options);

  onBeforeUnmount(() => plugin.remove(confirm));

  return {
    ask,
    anchorStyle: { anchorName: confirm.anchorVar },
  };
}
