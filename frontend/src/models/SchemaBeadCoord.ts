import type { IPoint } from './Point';

export type SchemaBeadCoord = `${number}:${number}`;

export function serializeSchemaBeadCoord(x: number, y: number): SchemaBeadCoord {
  return `${x}:${y}`;
}

export function parseSchemaBeadCoord(coord: string): IPoint {
  const [x, y] = coord.split(':').map(Number);
  return { x: x!, y: y! };
}
