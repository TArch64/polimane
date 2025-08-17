export type SchemaSizeDirection = 'top' | 'left' | 'right' | 'bottom';
export type SchemaSize = Record<SchemaSizeDirection, number>;

export type SchemaBeedCoord = `${number}:${number}`;

export function parseSchemaBeedCoord(coord: SchemaBeedCoord): [number, number] {
  return coord.split(':').map(Number) as [number, number];
}

export interface ISchema {
  id: string;
  name: string;
  palette: string[];
  createdAt: string;
  updatedAt: string;
  screenshotedAt: string | null;
  screenshotPath: string | null;
  size: SchemaSize;
  beads: Record<SchemaBeedCoord, string>;
}
