import { computed, reactive, type Ref, ref, watch, type WatchStopHandle } from 'vue';
import type { ISchema } from '@/models';
import { combineStopHandles } from '@/helpers';

const SAVE_TIMEOUT = 2000;

export interface IEditorSaveDispatcher {
  hasUnsavedChanges: boolean;
  enable: () => void;
  disable: () => void;
  flush: () => Promise<void>;
  abandon: () => void;
}

type EditorSaveCallback = (patch: Partial<ISchema>) => Promise<void>;
type WatchableAttribute = keyof Omit<ISchema, 'id'>;

export function useEditorSaveDispatcher(schema: Ref<ISchema>, onSave: EditorSaveCallback): IEditorSaveDispatcher {
  let saveTimeout: TimeoutId | null = null;
  let stopWatch: VoidFunction | null = null;
  const unsavedChanges = ref<Partial<ISchema> | null>(null);
  const hasUnsavedChanges = computed(() => !!unsavedChanges.value);

  async function dispatchSave(): Promise<void> {
    saveTimeout = null;
    if (unsavedChanges.value) {
      await onSave(unsavedChanges.value);
      unsavedChanges.value = null;
    }
  }

  function watchSavableAttribute(attr: WatchableAttribute): WatchStopHandle {
    return watch(() => schema.value[attr], (value) => {
      if (saveTimeout) {
        clearTimeout(saveTimeout);
      }

      unsavedChanges.value ??= {};
      // @ts-expect-error no easy way to match types
      unsavedChanges.value[attr] = value;
      saveTimeout = setTimeout(dispatchSave, SAVE_TIMEOUT);
    }, { deep: true });
  }

  function enable(): void {
    const attrStopWatchers: WatchStopHandle[] = [];

    for (const attr of Object.keys(schema.value)) {
      if (attr === 'id') {
        continue;
      }

      attrStopWatchers.push(watchSavableAttribute(attr as WatchableAttribute));
    }

    stopWatch = combineStopHandles(...attrStopWatchers);
  }

  const disable = () => stopWatch?.();

  async function flush(): Promise<void> {
    if (saveTimeout) {
      clearTimeout(saveTimeout);
    }

    if (unsavedChanges.value) {
      await dispatchSave();
    }
  }

  function abandon(): void {
    if (saveTimeout) {
      clearTimeout(saveTimeout);
    }

    unsavedChanges.value = null;
  }

  return reactive({
    hasUnsavedChanges,
    enable,
    disable,
    flush,
    abandon,
  });
}
