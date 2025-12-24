import { AccessLevel, BeadKind, Direction, SchemaLayout } from '@/enums';
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
  layout: SchemaLayout;
  screenshotedAt: string | null;
  screenshotPath: string | null;
  size: SchemaSize;
  beads: SchemaBeads;
}

export const SchemaUpdatableAttrs = [
  'name',
  'palette',
  'backgroundColor',
  'size',
  'beads',
] as const satisfies readonly (keyof ISchema)[];

export type SchemaUpdatableAttr = typeof SchemaUpdatableAttrs[number];
export type SchemaUpdate = Partial<Pick<ISchema, SchemaUpdatableAttr>>;

export function isSchemaUpdatableAttr(attr: string): attr is SchemaUpdatableAttr {
  return SchemaUpdatableAttrs.includes(attr as SchemaUpdatableAttr);
}
