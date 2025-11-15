import type { ListSchema } from '../homeStore';

export interface ISchemaCreateRequest {
  name: string;
}

export interface ISchemaCreateStrategy {
  do(request: ISchemaCreateRequest): Promise<ListSchema>;
}
