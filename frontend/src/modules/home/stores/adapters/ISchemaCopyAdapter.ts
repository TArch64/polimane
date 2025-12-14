import type { ListSchema } from '../homeStore';

export interface ISchemaCopyAdapter {
  do(schema: ListSchema): Promise<ListSchema>;
}
