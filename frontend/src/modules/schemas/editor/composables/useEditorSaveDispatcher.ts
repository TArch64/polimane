import { computed, reactive, type Ref, ref, watch, type WatchStopHandle } from 'vue';
import type { ISchema } from '@/models';

const SAVE_TIMEOUT = 2000;

export interface IEditorSaveDispatcher {
  hasUnsavedChanges: boolean;
  enable: () => void;
  disable: () => void;
  flush: () => Promise<void>;
}

export function useEditorSaveDispatcher(schema: Ref<ISchema>, onSave: () => Promise<void>): IEditorSaveDispatcher {
  let stopWatch: WatchStopHandle | null = null;
  const saveTimeout = ref<TimeoutId | null>(null);
  const hasUnsavedChanges = computed(() => !!saveTimeout.value);

  async function dispatchSave(): Promise<void> {
    saveTimeout.value = null;
    await onSave();
  }

  function onChange() {
    if (saveTimeout.value) {
      clearTimeout(saveTimeout.value);
    }

    saveTimeout.value = setTimeout(dispatchSave, SAVE_TIMEOUT);
  }

  function enable(): void {
    stopWatch = watch(schema, onChange, { deep: true });
  }

  function disable(): void {
    stopWatch?.();
  }

  async function flush(): Promise<void> {
    if (saveTimeout.value) {
      clearTimeout(saveTimeout.value);
      await dispatchSave();
    }
  }

  return reactive({
    hasUnsavedChanges,
    enable,
    disable,
    flush,
  });
}
