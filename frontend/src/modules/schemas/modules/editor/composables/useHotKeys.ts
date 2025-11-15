import { useEventListener } from '@vueuse/core';
import {
  computed,
  inject,
  type InjectionKey,
  type MaybeRefOrGetter,
  provide,
  reactive,
  toValue,
  watch,
} from 'vue';
import { newId } from '@/helpers';
import { isMac } from '@/config';

export type HotKeyExec = (event: KeyboardEvent) => void;
export type HotKeyDef = [expr: string, exec: HotKeyExec];
export type HotKeysDef = Record<HotKeyDef[0], HotKeyDef[1]>;
export type HotKeysVariantsDef = Record<'win' | 'mac', HotKeysDef>;
export type AnyHotKeysDef = HotKeysVariantsDef | HotKeysDef | HotKeyDef[];

export interface IHotKeysOptions {
  isActive?: MaybeRefOrGetter<boolean>;
}

interface IHotKey {
  meta: boolean;
  alt: boolean;
  ctrl: boolean;
  shift: boolean;
  key: string;
  expr: string;
  exec: HotKeyExec;
}

interface IHotKeysHandler {
  activate: (clientId: string, hotKeys: IHotKey[]) => void;
  deactivate: (clientId: string) => void;
}

const HOT_KEYS_HANDLER = Symbol('HotKeysHandler') as InjectionKey<IHotKeysHandler>;

export function provideHotKeysHandler(): void {
  const register = reactive(new Map<string, IHotKey[]>());

  function activate(clientId: string, hotKeys: IHotKey[]): void {
    register.set(clientId, hotKeys);
  }

  function deactivate(clientId: string): void {
    register.delete(clientId);
  }

  const isActive = computed(() => register.size > 0);
  const target = computed(() => isActive.value ? document.documentElement : null);

  useEventListener(target, 'keydown', (event) => {
    for (const hotKeys of register.values()) {
      const hotKey = hotKeys.find((hotKey) => {
        return hotKey.meta === event.metaKey
          && hotKey.shift === event.shiftKey
          && hotKey.alt === event.altKey
          && hotKey.ctrl === event.ctrlKey
          && hotKey.key === event.code;
      });

      if (hotKey) {
        event.preventDefault();
        hotKey.exec(event);
        return;
      }
    }
  }, { capture: true });

  provide(HOT_KEYS_HANDLER, {
    activate,
    deactivate,
  });
}

function normalizeDefs(def: AnyHotKeysDef): HotKeyDef[] {
  if (Array.isArray(def)) {
    return def;
  }

  if ('win' in def && 'mac' in def) {
    return isMac ? Object.entries(def.mac) : Object.entries(def.win);
  }

  return Object.entries(def);
}

export interface IHotKeysMeta {
  titles: Record<string, string>;
}

export function useHotKeys(def: AnyHotKeysDef, options: IHotKeysOptions = {}): IHotKeysMeta {
  const clientId = newId();
  const handler = inject(HOT_KEYS_HANDLER)!;

  const entries = normalizeDefs(def);
  const isActive = computed(() => toValue(options.isActive) ?? true);

  const hotKeys = entries.map(([expr, exec]): IHotKey => {
    const parts = expr.split('_');

    const hotKey: IHotKey = {
      meta: false,
      shift: false,
      alt: false,
      ctrl: false,
      key: '',
      expr,
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

  watch(isActive, (isActive) => {
    isActive
      ? handler.activate(clientId, hotKeys)
      : handler.deactivate(clientId);
  }, { immediate: true });

  const titles = computed(() => Object.fromEntries(
    hotKeys.map((hotKey) => [
      hotKey.expr,
      [
        hotKey.meta && (isMac ? '⌘' : 'Він'),
        hotKey.ctrl && (isMac ? '⌃' : 'Ктрл'),
        hotKey.alt && (isMac ? '⌥' : 'Альт'),
        hotKey.shift && (isMac ? '⇧' : 'Шфифт'),
        hotKey.key
          .replace('Key', '')
          .replace('Digit', '')
          .replace('=', hotKey.meta ? '+' : '='),
      ].filter(Boolean).join(' '),
    ]),
  ));

  return reactive({ titles });
}
