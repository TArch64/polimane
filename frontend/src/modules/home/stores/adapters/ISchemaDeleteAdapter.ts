import type { ListSchema } from '../homeStore';

export interface IDeleteManySchemasRequest {
  ids: string[];
}

export interface ISchemaDeleteAdapter {
  do(schema: ListSchema): Promise<void>;

  doMany(ids: string[]): Promise<void>;
}
