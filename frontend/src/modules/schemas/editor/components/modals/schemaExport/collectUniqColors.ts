import { getBeadSettings, type ISchema, isRefBead, type SchemaBead } from '@/models';
import type { BeadContentKind } from '@/enums';

export function collectUniqColors(schema: ISchema): string[] {
  const set = new Set<string>();

  for (const bead of Object.values(schema.beads)) {
    if (isRefBead(bead)) {
      continue;
    }

    const settings = getBeadSettings(bead as SchemaBead<BeadContentKind>);
    set.add(settings.color);
  }

  return Array.from(set);
}
