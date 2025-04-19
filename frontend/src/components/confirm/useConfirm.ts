import { onUnmounted } from 'vue';
import { type ConfirmCreateOptions, ConfirmPlugin } from './ConfirmPlugin';

export interface IConfirm {
  ask: () => Promise<boolean>;
  anchorStyle: { anchorName: string };
}

export function useConfirm(options: ConfirmCreateOptions): IConfirm {
  const plugin = ConfirmPlugin.inject();
  const confirm = plugin.create(options);
  const ask = () => confirm.ask();

  onUnmounted(() => plugin.remove(confirm));

  return { ask, anchorStyle: { anchorName: confirm.anchorVar } };
}
