import type { ISchemaPattern } from '@/models';
import { getMappedValue } from '@/helpers';
import { PatternType } from '@/enums';
import PatternAddSquareRowModal from './PatternAddSquareRowModal.vue';

export function getPatternAddRowModal(pattern: ISchemaPattern) {
  return getMappedValue(pattern.type, {
    [PatternType.SQUARE]: PatternAddSquareRowModal,
    [PatternType.DIAMOND]: PatternAddSquareRowModal,
  });
}

export { default as PatternAddModal } from './PatternAddModal.vue';
export { default as PatternRenameModal } from './PatternRenameModal.vue';
