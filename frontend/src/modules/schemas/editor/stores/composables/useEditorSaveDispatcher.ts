import { computed, reactive, type Ref, ref, watch, type WatchStopHandle } from 'vue';
import { combineStopHandles, getObjectKeys } from '@/helpers';
import type { ISchema } from '@/models';
import type { SafeAny } from '@/types';

const SAVE_TIMEOUT = 30_000;

export interface IEditorSaveDispatcher {
  hasUnsavedChanges: boolean;
  isSaving: boolean;
  enable: () => void;
  disable: () => void;
  flush: () => Promise<void>;
  abandon: () => void;
}

type EditorSaveCallback = (patch: Partial<ISchema>) => Promise<void>;

const NON_WATCHABLE_ATTRIBUTES = ['id', 'createdAt', 'updatedAt', 'screenshotedAt'] as const;

function isNonWatchableAttribute(attr: string): attr is NonWatchableAttribute {
  return NON_WATCHABLE_ATTRIBUTES.includes(attr as NonWatchableAttribute);
}

type NonWatchableAttribute = (typeof NON_WATCHABLE_ATTRIBUTES)[number];
type WatchableAttribute = keyof Omit<ISchema, NonWatchableAttribute>;

export function useEditorSaveDispatcher(schema: Ref<ISchema>, onSave: EditorSaveCallback): IEditorSaveDispatcher {
  let saveTimeout: TimeoutId | null = null;
  let stopWatch: VoidFunction | null = null;
  const unsavedChanges = ref<Partial<ISchema> | null>(null);
  const hasUnsavedChanges = computed(() => !!unsavedChanges.value);
  const isSaving = ref(false);

  async function dispatchSave(): Promise<void> {
    saveTimeout = null;

    if (unsavedChanges.value) {
      try {
        isSaving.value = true;
        await onSave(unsavedChanges.value);
        unsavedChanges.value = null;
      } finally {
        isSaving.value = false;
      }
    }
  }

  function watchSavableAttribute(attr: WatchableAttribute): WatchStopHandle {
    return watch(() => schema.value[attr], (value) => {
      if (saveTimeout) {
        clearTimeout(saveTimeout);
      }

      unsavedChanges.value ??= {};
      (unsavedChanges.value as Record<WatchableAttribute, SafeAny>)[attr] = value;
      saveTimeout = setTimeout(dispatchSave, SAVE_TIMEOUT);
    }, { deep: true });
  }

  function enable(): void {
    const attrStopWatchers: WatchStopHandle[] = [];

    for (const attr of getObjectKeys(schema.value)) {
      if (isNonWatchableAttribute(attr)) {
        continue;
      }

      attrStopWatchers.push(watchSavableAttribute(attr));
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
    isSaving,
    enable,
    disable,
    flush,
    abandon,
  });
}
