import type { ListSchema } from '../homeStore';

export interface ISchemaCopyStrategy {
  do(schema: ListSchema): Promise<ListSchema>;
}
