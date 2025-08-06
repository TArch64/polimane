import { type FunctionalComponent, h } from 'vue';
import type { ISchemaPattern, ISchemaRow } from '@/models';
import { getMappedValue } from '@/helpers';
import { PatternType } from '@/enums';
import CanvasSquareRow from './CanvasSquareRow.vue';

export interface ICanvasRowProps {
  pattern: ISchemaPattern;
  row: ISchemaRow;
}

export const CanvasRow: FunctionalComponent<ICanvasRowProps> = (props) => getMappedValue(props.pattern.type, {
  [PatternType.SQUARE]: () => h(CanvasSquareRow, props),
  [PatternType.DIAMOND]: () => h(CanvasSquareRow, props),
});
