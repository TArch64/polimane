import { type ConfirmCreateOptions, ConfirmPlugin } from './ConfirmPlugin';

export interface IConfirm {
  ask: () => Promise<boolean>;
}

export function useConfirm(options: ConfirmCreateOptions): IConfirm {
  const plugin = ConfirmPlugin.inject();
  const confirm = plugin.create(options);
  const ask = () => confirm.ask();
  return { ask };
}
