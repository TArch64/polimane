import type { ISchemaPattern } from './ISchemaPattern';

export interface ISchemaContent {
  patterns: ISchemaPattern[];
}

export interface ISchema {
  id: string;
  name: string;
  content: ISchemaContent;
}
