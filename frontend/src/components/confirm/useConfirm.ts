import { onUnmounted } from 'vue';
import { useTopElement } from '@/composables';
import { type ConfirmCreateInternalOptions, ConfirmPlugin } from './ConfirmPlugin';
import type { IConfirmAskOptions } from './Confirm';

export interface IConfirm {
  ask: (options?: IConfirmAskOptions) => Promise<boolean>;
  anchorStyle: { anchorName: string };
}

export type ConfirmCreateOptions = Omit<ConfirmCreateInternalOptions, 'getTopEl'>;

export function useConfirm(options: ConfirmCreateOptions): IConfirm {
  const plugin = ConfirmPlugin.inject();
  const topEl = useTopElement();

  const confirm = plugin.create({
    ...options,
    getTopEl: () => topEl.value,
  });

  const ask = (options: IConfirmAskOptions = {}) => confirm.ask(options);

  onUnmounted(() => plugin.remove(confirm));

  return {
    ask,
    anchorStyle: { anchorName: confirm.anchorVar },
  };
}
