export type SchemaSizeDirection = 'top' | 'left' | 'right' | 'bottom';
export type SchemaSize = Record<SchemaSizeDirection, number>;

export type SchemaBeadCoord = `${number}:${number}`;

export function serializeSchemaBeadCoord(x: number, y: number): SchemaBeadCoord {
  return `${x}:${y}`;
}

export function parseSchemaBeadCoord(coord: SchemaBeadCoord): [number, number] {
  return coord.split(':').map(Number) as [number, number];
}

export interface ISchema {
  id: string;
  name: string;
  palette: string[];
  createdAt: string;
  updatedAt: string;
  backgroundColor: string;
  screenshotedAt: string | null;
  screenshotPath: string | null;
  size: SchemaSize;
  beads: Record<SchemaBeadCoord, string>;
}
