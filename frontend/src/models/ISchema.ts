import type { ISchemaPattern } from './ISchemaPattern';

export interface ISchema {
  id: string;
  name: string;
  palette: string[];
  createdAt: string;
  updatedAt: string;
  content: ISchemaPattern[];
}
