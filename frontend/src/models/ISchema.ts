import type { ISchemaPattern } from './ISchemaPattern';

export interface ISchema {
  id: string;
  name: string;
  content: ISchemaPattern[];
}
