import { type BeadContentKind, BeadKind, type BeadSpannableKind } from '@/enums';
import type { IPoint } from './Point';
import type { BeadCoord } from './SchemaBeadCoord';

export type SchemaBeadMap = {
  [BeadKind.CIRCLE]: ISchemaCircleBead;
  [BeadKind.BUGLE]: ISchemaBugleBead;
  [BeadKind.REF]: ISchemaRefBead;
};

export type SchemaBead<K extends BeadKind = BeadKind> = Partial<Pick<SchemaBeadMap, K>> & {
  kind: K;
};

export type SchemaContentBead = SchemaBead<BeadContentKind>;
export type SchemaSpannableBead = SchemaBead<BeadSpannableKind>;

export interface ISchemaBaseBead {
  color: string;
}

export interface ISchemaCircleBead extends ISchemaBaseBead {
}

export interface ISchemaBugleBead extends ISchemaBaseBead {
  span: IPoint;
}

export interface ISchemaRefBead {
  to: BeadCoord;
}

export function getBeadSettings<K extends BeadKind = BeadKind>(bead: SchemaBead<K>): SchemaBeadMap[K] {
  return bead[bead.kind as K] as SchemaBeadMap[K];
}

export function isRefBead(bead: SchemaBead): boolean {
  return bead.kind === BeadKind.REF;
}
