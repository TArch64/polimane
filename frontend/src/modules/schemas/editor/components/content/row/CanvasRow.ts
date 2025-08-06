import { type FunctionalComponent, h } from 'vue';
import type { ISchemaPattern, SchemaRow } from '@/models';
import { getMappedValue } from '@/helpers';
import { PatternType } from '@/enums';
import CanvasSquareRow from './CanvasSquareRow.vue';

export interface ICanvasRowProps {
  pattern: ISchemaPattern;
  row: SchemaRow;
}

export const CanvasRow: FunctionalComponent<ICanvasRowProps> = (props) => getMappedValue(props.pattern.type, {
  [PatternType.SQUARE]: () => h(CanvasSquareRow, props),
  [PatternType.DIAMOND]: () => h(CanvasSquareRow, props),
});
