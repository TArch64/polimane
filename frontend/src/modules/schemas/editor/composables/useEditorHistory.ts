import { computed, reactive, ref, type Ref, shallowRef } from 'vue';
import { watchDebounced } from '@vueuse/core';
import type { ISchema } from '@/models';
import { setSchemaRelations } from '@/modules/schemas/editor/models';
import { compressObject, decompressObject } from '@/helpers';

export interface IEditorHistory {
  init: () => Promise<void>;
  destroy: () => void;
  undo: () => Promise<void>;
  redo: () => Promise<void>;
}

export function useEditorHistory(schema: Ref<ISchema>): IEditorHistory {
  let stopWatcher: VoidFunction;

  const history = shallowRef<Blob[]>([]);
  const cursor = ref(-1);

  const canUndo = computed(() => cursor.value > 0);
  const canRedo = computed(() => cursor.value < history.value.length - 1);

  async function commit() {
    if (canRedo.value) {
      history.value = history.value.slice(0, cursor.value + 1);
    }

    const blob = await compressObject(schema.value.content);
    history.value = [...history.value, blob].slice(-30);
    cursor.value = history.value.length - 1;
  }

  function startWatcher(): void {
    stopWatcher = watchDebounced(() => schema.value.content, commit, {
      deep: true,
      debounce: 200,
    });
  }

  async function init(): Promise<void> {
    history.value = [];
    await commit();
    startWatcher();
  }

  function destroy(): void {
    stopWatcher();
    history.value = [];
  }

  async function restoreVersion(shift: number): Promise<void> {
    if (cursor.value + shift < 0 || cursor.value + shift >= history.value.length) {
      return;
    }

    stopWatcher();
    cursor.value += shift;
    const compressed = history.value[cursor.value];

    const newSchema = {
      ...schema.value,
      content: await decompressObject<ISchema['content']>(compressed),
    };

    setSchemaRelations(newSchema);
    schema.value = newSchema;
    startWatcher();
  }

  return reactive({
    init,
    destroy,
    undo: async () => {
      canUndo.value && await restoreVersion(-1);
    },
    redo: async () => {
      canRedo.value && await restoreVersion(1);
    },
  });
}
