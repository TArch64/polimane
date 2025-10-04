export type SchemaSizeDirection = 'top' | 'left' | 'right' | 'bottom';
export type SchemaSize = Record<SchemaSizeDirection, number>;

export type SchemaBeadCoord = `${number}:${number}`;
export type SchemaBeadCoordTuple = [x: number, y: number];
export type SchemaBeads = Record<SchemaBeadCoord, string>;

export function serializeSchemaBeadCoord(x: number, y: number): SchemaBeadCoord {
  return `${x}:${y}`;
}

export function parseSchemaBeadCoord(coord: SchemaBeadCoord): SchemaBeadCoordTuple {
  return coord.split(':').map(Number) as SchemaBeadCoordTuple;
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
  beads: SchemaBeads;
}
