export type SchemaContentSize = Record<'top' | 'left' | 'right' | 'bottom', number>;
export type SchemaBeedCoordinate = `${number}:${number}`;

export interface ISchemaContent {
  size: SchemaContentSize;
  beads: Record<SchemaBeedCoordinate, string>;
}

export interface ISchema {
  id: string;
  name: string;
  palette: string[];
  createdAt: string;
  updatedAt: string;
  screenshotedAt: string | null;
  screenshotPath: string | null;
  content: ISchemaContent;
}
