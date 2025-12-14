import type { SchemaUpdate } from '@/models';
import type { ListSchema } from '../homeStore';

export interface ISchemaUpdateAdapter {
  do(schema: ListSchema, update: SchemaUpdate): Promise<void>;
}
