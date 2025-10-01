import type { ISchema } from '@/models';

export function collectUniqColors(schema: ISchema): string[] {
  const list = Object.values(schema.beads);
  return Array.from(new Set(list));
}
