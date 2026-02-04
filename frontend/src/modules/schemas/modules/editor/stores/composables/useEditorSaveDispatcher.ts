import { computed, reactive, type Ref, ref, watch, type WatchStopHandle } from 'vue';
import { combineStopHandles, getObjectKeys } from '@/helpers';
import {
  type ISchema,
  isSchemaUpdatableAttr,
  type SchemaUpdatableAttr,
  type SchemaUpdate,
} from '@/models';
import type { MaybePromise } from '@/types';

const SAVE_TIMEOUT = 30_000;

export interface IEditorSaveDispatcher {
  hasUnsavedChanges: boolean;
  isSaving: boolean;
  enable: () => void;
  disable: () => void;
  flush: () => Promise<void>;
  abandon: () => void;
}

type ChangeCallbackMap = {
  [K in SchemaUpdatableAttr]: (value: Required<SchemaUpdate>[K]) => void;
};

export interface IEditorSaveDispatcherOptions {
  onSave: (patch: SchemaUpdate) => MaybePromise<void>;
  onChange: Partial<ChangeCallbackMap>;
}

export function useEditorSaveDispatcher(schema: Ref<ISchema>, options: IEditorSaveDispatcherOptions): IEditorSaveDispatcher {
  let saveTimeout: TimeoutId | null = null;
  let stopWatch: VoidFunction | null = null;
  const unsavedChanges = ref<SchemaUpdate | null>(null);
  const hasUnsavedChanges = computed(() => !!unsavedChanges.value);
  const isSaving = ref(false);

  async function dispatchSave(): Promise<void> {
    saveTimeout = null;

    if (unsavedChanges.value) {
      try {
        isSaving.value = true;
        await options.onSave(unsavedChanges.value);
        unsavedChanges.value = null;
      } finally {
        isSaving.value = false;
      }
    }
  }

  function watchSavableAttribute<A extends SchemaUpdatableAttr>(attr: A): WatchStopHandle {
    return watch(() => schema.value[attr], (value) => {
      if (saveTimeout) {
        clearTimeout(saveTimeout);
      }

      unsavedChanges.value ??= {};
      unsavedChanges.value[attr] = value;
      saveTimeout = setTimeout(dispatchSave, SAVE_TIMEOUT);
      options.onChange[attr]?.(value);
    }, { deep: true });
  }

  function enable(): void {
    const attrStopWatchers: WatchStopHandle[] = [];

    for (const attr of getObjectKeys(schema.value)) {
      if (!isSchemaUpdatableAttr(attr)) {
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
