import type { ISchemaPattern } from '@/models';
import { getMappedValue } from '@/helpers';
import { PatternType } from '@/enums';
import EditorAddSquareRowModal from './EditorAddSquareRowModal.vue';

export function getPatternAddRowModal(pattern: ISchemaPattern) {
  return getMappedValue(pattern.type, {
    [PatternType.SQUARE]: EditorAddSquareRowModal,
    [PatternType.DIAMOND]: EditorAddSquareRowModal,
  });
}

export { default as EditorAddPatternModal } from './EditorAddPatternModal.vue';
export { default as PatternRenameModal } from './PatternRenameModal.vue';
