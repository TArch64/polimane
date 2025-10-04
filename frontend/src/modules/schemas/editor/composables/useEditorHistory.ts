import { computed, markRaw, reactive, ref, type Ref, shallowRef } from 'vue';
import { watchDebounced } from '@vueuse/core';
import type { ISchema } from '@/models';
import { ObjectCompressor } from '@/helpers';

type HistoryRecord = Pick<ISchema, 'beads' | 'size'>;

export interface IEditorHistory {
  canUndo: boolean;
  canRedo: boolean;
  init: () => Promise<void>;
  destroy: () => void;
  undo: () => Promise<void>;
  redo: () => Promise<void>;
}

export function useEditorHistory(schema: Ref<ISchema>): IEditorHistory {
  let stopWatcher: VoidFunction;
  const objectCompressor = ObjectCompressor.typed<HistoryRecord>();

  const history = shallowRef<Blob[]>([]);
  const cursor = ref(-1);

  const canUndo = computed(() => cursor.value > 0);
  const canRedo = computed(() => cursor.value < history.value.length - 1);

  async function commit() {
    if (canRedo.value) {
      history.value = history.value.slice(0, cursor.value + 1);
    }

    const blob = markRaw(await objectCompressor.compress({
      beads: schema.value.beads,
      size: schema.value.size,
    }));

    history.value = [...history.value, blob].slice(-30);
    cursor.value = history.value.length - 1;
  }

  function startWatcher(): void {
    stopWatcher = watchDebounced([
      () => schema.value.beads,
      () => schema.value.size,
    ], commit, {
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
    const compressed = history.value[cursor.value]!;

    schema.value = {
      ...schema.value,
      ...await objectCompressor.decompress(compressed),
    };

    startWatcher();
  }

  async function undo(): Promise<void> {
    if (canUndo.value) await restoreVersion(-1);
  }

  async function redo(): Promise<void> {
    if (canRedo.value) await restoreVersion(1);
  }

  return reactive({
    init,
    destroy,
    undo,
    canUndo,
    redo,
    canRedo,
  });
}
