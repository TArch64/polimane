import type { ISchemaPattern } from './ISchemaObject';

export interface ISchema {
  id: string;
  name: string;
  content: ISchemaPattern[];
}
