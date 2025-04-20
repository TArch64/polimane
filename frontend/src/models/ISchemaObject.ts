import { PatternKind, SchemaObjectType } from '@/enums';

export interface ISchemaObject {
  id: string;
  type: SchemaObjectType;
}

export interface ISchemaPattern extends ISchemaObject {
  name: string;
  type: SchemaObjectType.PATTERN;
  kind: PatternKind;
  content: ISchemaObject[];
}
