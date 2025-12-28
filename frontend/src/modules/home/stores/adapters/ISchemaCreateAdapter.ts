import { SchemaLayout } from '@/enums';
import type { ListSchema } from '../homeStore';

export interface ISchemaCreateRequest {
  name: string;
  layout: SchemaLayout;
  folderId?: string;
}

export interface ISchemaCreateAdapter {
  do(request: ISchemaCreateRequest): Promise<ListSchema>;
}
