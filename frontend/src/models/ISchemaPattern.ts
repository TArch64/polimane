import { PatternType } from '@/enums';
import type { ISchemaWithContent } from './ISchemaWithContent';
import type { ISchemaRow } from './ISchemaRow';

export interface ISchemaPattern extends ISchemaWithContent<ISchemaRow> {
  name: string;
  type: PatternType;
}
