import { Direction } from '@/enums';
import type { SchemaBead } from './ISchemaBead';
import type { SchemaBeadCoord } from './SchemaBeadCoord';

export type SchemaSize = Record<Direction, number>;
export type SchemaBeads = Record<SchemaBeadCoord, SchemaBead>;

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
