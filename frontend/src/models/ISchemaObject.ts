import { PatternType } from '@/enums';

export interface ISchemaObject {
  id: string;
}

export interface ISchemaPattern extends ISchemaObject {
  name: string;
  type: PatternType;
  content: ISchemaRow[];
}

export interface ISchemaRow extends ISchemaObject {
  content: ISchemaObject[];
}
