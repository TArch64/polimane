import { SchemaLayout } from '@/enums';
import type { ListSchema } from '../homeStore';

export interface ISchemaCreateRequest {
  name: string;
  layout: SchemaLayout;
}

export interface ISchemaCreateAdapter {
  do(request: ISchemaCreateRequest): Promise<ListSchema>;
}
