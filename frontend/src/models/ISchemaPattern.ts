import { PatternType } from '@/enums';
import type { ISchemaWithContent } from './ISchemaWithContent';
import type { ISchemaRow } from './ISchemaRow';

interface ISchemaPatternBase extends ISchemaWithContent<ISchemaRow> {
  name: string;
}

export interface ISchemaPatternSquareAttrs {
  size: number;
}

export interface ISchemaPatternSquare extends ISchemaPatternBase {
  type: PatternType.SQUARE;
  square: ISchemaPatternSquareAttrs;
}

export interface ISchemaPatternDiamondAttrs {
  size: number;
  sideSize: number;
}

export interface ISchemaPatternDiamond extends ISchemaPatternBase {
  type: PatternType.DIAMOND;
  diamond: ISchemaPatternDiamondAttrs;
}

export type SchemaPattern = ISchemaPatternSquare | ISchemaPatternDiamond;
