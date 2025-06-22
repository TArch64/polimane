import { type FunctionalComponent, h } from 'vue';
import type { ISchemaPattern } from '@/models';
import { getMappedValue } from '@/helpers';
import { PatternType } from '@/enums';
import RowAddSquareModal from './RowAddSquareModal.vue';

interface IRowAddModalProps {
  pattern: ISchemaPattern;
  toIndex?: number;
}

const getComponent = (pattern: ISchemaPattern) => getMappedValue(pattern.type, {
  [PatternType.SQUARE]: RowAddSquareModal,
  [PatternType.DIAMOND]: RowAddSquareModal,
});

export const RowAddModal: FunctionalComponent<IRowAddModalProps> = (props) => {
  return h(getComponent(props.pattern), props);
};
