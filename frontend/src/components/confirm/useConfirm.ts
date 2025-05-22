import { onUnmounted } from 'vue';
import { type ConfirmCreateOptions, ConfirmPlugin } from './ConfirmPlugin';
import type { IConfirmAskOptions } from './Confirm';

export interface IConfirm {
  ask: (options?: IConfirmAskOptions) => Promise<boolean>;
  anchorStyle: { anchorName: string };
}

export function useConfirm(options: ConfirmCreateOptions): IConfirm {
  const plugin = ConfirmPlugin.inject();
  const confirm = plugin.create(options);
  const ask = (options: IConfirmAskOptions = {}) => confirm.ask(options);

  onUnmounted(() => plugin.remove(confirm));

  return { ask, anchorStyle: { anchorName: confirm.anchorVar } };
}
