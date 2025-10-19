import type {
  ISchemaBugleBead,
  ISchemaCircleBead,
  SchemaBead,
  SchemaBeadCoord,
  SchemaBeadMap,
} from '@/models';
import { type BeadContentKind, BeadKind } from '@/enums';

type BeadBuilders = {
  [k in BeadContentKind]: (color: string) => SchemaBeadMap[k];
};

export function useBeadFactory() {
  const buildCircle = (color: string): ISchemaCircleBead => ({ color });

  const buildBugle = (color: string): ISchemaBugleBead => ({
    color,
    span: { x: 0, y: 0 },
  });

  const builders: BeadBuilders = {
    [BeadKind.CIRCLE]: buildCircle,
    [BeadKind.BUGLE]: buildBugle,
  };

  function create(kind: BeadContentKind, color: string | null): SchemaBead | null {
    if (!color) {
      return null;
    }
    return {
      kind,
      [kind]: builders[kind](color),
    };
  }

  const createRef = (to: SchemaBeadCoord): SchemaBead => ({
    kind: BeadKind.REF,
    ref: { to },
  });

  return { create, createRef };
}
