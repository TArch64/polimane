import { getBeadSettings, type ISchema, isRefBead, type SchemaContentBead } from '@/models';

export function collectUniqColors(schema: ISchema): string[] {
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
