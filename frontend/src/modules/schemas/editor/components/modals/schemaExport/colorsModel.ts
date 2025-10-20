import { getBeadSettings, type ISchema, isRefBead, type SchemaContentBead } from '@/models';

export interface ISchemaColorModel {
  initial: string;
  current: string;
}

function collectUniqColors(schema: ISchema): string[] {
  const set = new Set<string>();

  for (const bead of Object.values(schema.beads)) {
    if (isRefBead(bead)) {
      continue;
    }

    const settings = getBeadSettings(bead as SchemaContentBead);
    set.add(settings.color);
  }

  return Array.from(set);
}

export function buildColorsModel(schema: ISchema): ISchemaColorModel[] {
  return collectUniqColors(schema).map((color) => ({
    initial: color,
    current: color,
  }));
}
