import { useEventListener } from '@vueuse/core';

export type HotKeyExec = () => void;
export type HotKeyDef = [expr: string, exec: HotKeyExec];
export type HotKeysDef = Record<HotKeyDef[0], HotKeyDef[1]>;

interface IHotKey {
  meta: boolean;
  shift: boolean;
  key: string;
  exec: HotKeyExec;
}

export function useHotKeys(def: HotKeysDef): void;
export function useHotKeys(def: HotKeyDef[]): void;
export function useHotKeys(def: HotKeysDef | HotKeyDef[]): void {
  const entries = Array.isArray(def) ? def : Object.entries(def);

  const hotKeys = entries.map(([expr, exec]): IHotKey => {
    const parts = expr.toLowerCase().split('_');

    const hotKey: IHotKey = {
      meta: false,
      shift: false,
      key: '',
      exec,
    };

    let parsing = parts.shift();

    while (parsing) {
      if (parsing in hotKey) {
        // @ts-expect-error -- dynamic key
        hotKey[parsing] = true;
      } else {
        hotKey.key = parsing;
      }

      parsing = parts.shift();
    }

    return hotKey;
  });

  useEventListener('keydown', (event) => {
    const key = event.key.toLowerCase();

    const hotKey = hotKeys.find((hotKey) => {
      return hotKey.meta === event.metaKey
        && hotKey.shift === event.shiftKey
        && hotKey.key === key;
    });

    if (hotKey) {
      event.preventDefault();
      hotKey.exec();
    }
  });
}
