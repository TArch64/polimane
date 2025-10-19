import { type BeadContentKind, BeadKind } from '@/enums';
import type { IPoint } from './Point';
import type { SchemaBeadCoord } from './SchemaBeadCoord';

export type SchemaBeadMap = {
  [BeadKind.CIRCLE]: ISchemaCircleBead;
  [BeadKind.BUGLE]: ISchemaBugleBead;
  [BeadKind.REF]: ISchemaRefBead;
};

export type SchemaBead<K extends BeadKind = BeadKind> = Partial<Pick<SchemaBeadMap, K>> & {
  kind: K;
};

export type SchemaContentBead = SchemaBead<BeadContentKind>;

export interface ISchemaBaseBead {
  color: string;
}

export interface ISchemaCircleBead extends ISchemaBaseBead {
}

export interface ISchemaBugleBead extends ISchemaBaseBead {
  span: IPoint;
}

export interface ISchemaRefBead {
  to: SchemaBeadCoord;
}

export function getBeadSettings<K extends BeadKind = BeadKind>(bead: SchemaBead<K>): SchemaBeadMap[K] {
  return bead[bead.kind as K] as SchemaBeadMap[K];
}

export function isRefBead(bead: SchemaBead): boolean {
  return bead.kind === BeadKind.REF;
}
