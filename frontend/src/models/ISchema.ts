import { AccessLevel, BeadKind, Direction } from '@/enums';
import type { SchemaBead } from './ISchemaBead';
import type { BeadCoord } from './SchemaBeadCoord';

export type SchemaSize = Record<Direction, number>;
export type SchemaBeads<K extends BeadKind = BeadKind> = Record<BeadCoord, SchemaBead<K>>;

export interface ISchema {
  id: string;
  name: string;
  palette: string[];
  createdAt: string;
  updatedAt: string;
  access: AccessLevel;
  backgroundColor: string;
  screenshotedAt: string | null;
  screenshotPath: string | null;
  size: SchemaSize;
  beads: SchemaBeads;
}

export type SchemaUpdate = Partial<Omit<ISchema, 'id' | 'updatedAt' | 'createdAt' | 'screenshotedAt' | 'screenshotPath' | 'access'>>;
