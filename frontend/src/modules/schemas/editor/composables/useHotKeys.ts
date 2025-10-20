import { useEventListener } from '@vueuse/core';
import { computed, type MaybeRefOrGetter, toValue } from 'vue';

export type HotKeyExec = () => void;
export type HotKeyDef = [expr: string, exec: HotKeyExec];
export type HotKeysDef = Record<HotKeyDef[0], HotKeyDef[1]>;

export interface IHotKeysOptions {
  isActive?: MaybeRefOrGetter<boolean>;
}

interface IHotKey {
  meta: boolean;
  shift: boolean;
  key: string;
  exec: HotKeyExec;
}

export function useHotKeys(def: HotKeysDef | HotKeyDef[], options: IHotKeysOptions = {}): void {
  const entries = Array.isArray(def) ? def : Object.entries(def);
  const isActive = computed(() => toValue(options.isActive) ?? true);

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

  const target = computed(() => isActive.value ? document.documentElement : null);

  useEventListener(target, 'keydown', (event) => {
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
