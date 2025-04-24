import type { ISchemaPattern } from '@/models';
import { getMapValue } from '@/helpers';
import { PatternType } from '@/enums';

import EditorAddSquareRowModal from './EditorAddSquareRowModal.vue';

export function getPatternAddRowModal(pattern: ISchemaPattern) {
  return getMapValue(pattern.type, {
    [PatternType.SQUARE]: EditorAddSquareRowModal,
    [PatternType.DIAMOND]: EditorAddSquareRowModal,
  });
}

export { default as EditorAddPatternModal } from './EditorAddPatternModal.vue';
