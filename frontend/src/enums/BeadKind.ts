import { getMappedValue } from '@/helpers';

export enum BeadKind {
  CIRCLE = 'circle',
}

export const BeadKindList = [
  BeadKind.CIRCLE,
] as const;

export function getBeadKindTitle(kind: BeadKind): string {
  return getMappedValue(kind, {
    [BeadKind.CIRCLE]: 'Кругла',
  });
}
