import { useModal } from '@/components/modal';
import { usePatternsStore } from '@/modules/schemas/editor/stores';
import { RowAddModal } from '../row';
import PatternAddModal from './PatternAddModal.vue';
import type { IAddingPattern } from './IAddingPattern';

export { default as PatternRenameModal } from './PatternRenameModal.vue';

export function usePatternAddModal() {
  const patternsStore = usePatternsStore();

  const addPatternModal = useModal<typeof PatternAddModal, IAddingPattern>(PatternAddModal);
  const addRowModal = useModal(RowAddModal);

  addPatternModal.onResult(async ({ pattern, toIndex }) => {
    await addRowModal.open({ pattern });
    patternsStore.patterns.insert(pattern, { toIndex });
  });

  return addPatternModal;
}
