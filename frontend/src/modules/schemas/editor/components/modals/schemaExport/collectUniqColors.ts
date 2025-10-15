import type { ISchema } from '@/models';

export function collectUniqColors(schema: ISchema): string[] {
  const list = Object.values(schema.beads).map((bead) => bead.color);
  return Array.from(new Set(list));
}
