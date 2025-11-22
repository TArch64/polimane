import type { ListSchema } from '../homeStore';

export interface ISchemaDeleteStrategy {
  do(schema: ListSchema): Promise<void>;
}
