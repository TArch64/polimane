import { PatternType } from '@/enums';
import type { ISchemaWithContent } from './ISchemaWithContent';
import type { SchemaRow } from './ISchemaRow';

export interface ISchemaPattern extends ISchemaWithContent<SchemaRow> {
  name: string;
  type: PatternType;
}
