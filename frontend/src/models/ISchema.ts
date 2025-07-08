import type { ISchemaPattern } from './ISchemaPattern';

export interface ISchema {
  id: number;
  name: string;
  palette: string[];
  content: ISchemaPattern[];
}
