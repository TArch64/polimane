import { defineStore } from 'pinia';
import { computed, markRaw, onScopeDispose, type Ref, ref, shallowRef } from 'vue';
import { watchDebounced } from '@vueuse/core';
import { ObjectCompressor } from '@/helpers';
import type { ISchema } from '@/models';

type HistoryRecord = Pick<ISchema, 'beads' | 'size' | 'backgroundColor'>;

export const useHistoryStore = defineStore('schemas/editor/history', () => {
  let stopWatcher: VoidFunction;
  const objectCompressor = ObjectCompressor.typed<HistoryRecord>();

  let schema: Ref<ISchema>;
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
      backgroundColor: schema.value.backgroundColor,
    }));

    history.value = [...history.value, blob].slice(-30);
    cursor.value = history.value.length - 1;
  }

  function startWatcher(): void {
    stopWatcher = watchDebounced([
      () => schema.value.beads,
      () => schema.value.size,
      () => schema.value.backgroundColor,
    ], commit, {
      deep: true,
      debounce: 200,
    });
  }

  async function init(schema_: Ref<ISchema>): Promise<void> {
    schema = schema_;
    history.value = [];
    await commit();
    startWatcher();
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

  onScopeDispose(() => {
    stopWatcher();
    history.value = [];
  });

  return {
    init,
    undo,
    canUndo,
    redo,
    canRedo,
  };
});
