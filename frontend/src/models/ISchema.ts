import { Direction } from '@/enums';
import type { IPoint } from './Point';

export type SchemaSize = Record<Direction, number>;

export type SchemaBeadCoord = `${number}:${number}`;
export type SchemaBeads = Record<SchemaBeadCoord, string>;

export function serializeSchemaBeadCoord(x: number, y: number): SchemaBeadCoord {
  return `${x}:${y}`;
}

export function parseSchemaBeadCoord(coord: string): IPoint {
  const [x, y] = coord.split(':').map(Number);
  return { x: x!, y: y! };
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
