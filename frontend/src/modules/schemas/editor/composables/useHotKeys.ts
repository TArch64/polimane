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
  alt: boolean;
  ctrl: boolean;
  shift: boolean;
  key: string;
  exec: HotKeyExec;
}

export function useHotKeys(def: HotKeysDef | HotKeyDef[], options: IHotKeysOptions = {}): void {
  const entries = Array.isArray(def) ? def : Object.entries(def);
  const isActive = computed(() => toValue(options.isActive) ?? true);

  const hotKeys = entries.map(([expr, exec]): IHotKey => {
    const parts = expr.split('_');

    const hotKey: IHotKey = {
      meta: false,
      shift: false,
      alt: false,
      ctrl: false,
      key: '',
      exec,
    };

    function next() {
      return parts.shift();
    }

    let parsing = next();

    while (parsing) {
      const modifier = parsing.toLowerCase();

      if (modifier in hotKey) {
        // @ts-expect-error -- dynamic key
        hotKey[modifier] = true;
      } else {
        hotKey.key = parsing;
      }

      parsing = next();
    }

    return hotKey;
  });

  const target = computed(() => isActive.value ? document.documentElement : null);

  useEventListener(target, 'keydown', (event) => {
    const hotKey = hotKeys.find((hotKey) => {
      return hotKey.meta === event.metaKey
        && hotKey.shift === event.shiftKey
        && hotKey.alt === event.altKey
        && hotKey.ctrl === event.ctrlKey
        && hotKey.key === event.code;
    });

    if (hotKey) {
      event.preventDefault();
      hotKey.exec();
    }
  }, { capture: true });
}
