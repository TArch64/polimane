import type { ISchemaWithContent } from './ISchemaWithContent';
import type { ISchemaBead } from './ISchemaBead';

interface ISchemaRowBase extends ISchemaWithContent<ISchemaBead> {
}

export interface ISchemaSquareRowAttrs {
  size: number;
}

export interface ISchemaSquareRow extends ISchemaRowBase {
  square: ISchemaSquareRowAttrs;
}

export interface ISchemaDiamondRowAttrs {
  size: number;
  sideSize: number;
}

export interface ISchemaDiamondRow extends ISchemaRowBase {
  diamond: ISchemaDiamondRowAttrs;
}

export type SchemaRow = ISchemaSquareRow | ISchemaDiamondRow;
