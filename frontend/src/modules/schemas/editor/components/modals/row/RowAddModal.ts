import { type FunctionalComponent, h } from 'vue';
import type { SchemaPattern } from '@/models';
import { getMappedValue } from '@/helpers';
import { PatternType } from '@/enums';
import RowAddSquareModal from './RowAddSquareModal.vue';
import RowAddDiamondModal from './RowAddDiamondModal.vue';

interface IRowAddModalProps {
  pattern: SchemaPattern;
  toIndex?: number;
}

const getComponent = (pattern: SchemaPattern) => getMappedValue(pattern.type, {
  [PatternType.SQUARE]: RowAddSquareModal,
  [PatternType.DIAMOND]: RowAddDiamondModal,
});

export const RowAddModal: FunctionalComponent<IRowAddModalProps> = (props) => {
  return h(getComponent(props.pattern), props);
};
