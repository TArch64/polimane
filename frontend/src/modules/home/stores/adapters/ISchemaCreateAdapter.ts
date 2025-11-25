import type { ListSchema } from '../homeStore';

export interface ISchemaCreateRequest {
  name: string;
}

export interface ISchemaCreateAdapter {
  do(request: ISchemaCreateRequest): Promise<ListSchema>;
}
