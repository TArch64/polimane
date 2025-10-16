import { getMappedValue } from '@/helpers';

export enum BeadKind {
  CIRCLE = 'circle',
  BUGLE = 'bugle',
}

export const BeadKindList = [
  BeadKind.CIRCLE,
  BeadKind.BUGLE,
] as const;

export const BeadSpannableList = [BeadKind.BUGLE] as const;
export type BeadSpannableKind = typeof BeadSpannableList[number];

export function isBeadSpannableKind(kind: BeadKind): kind is BeadSpannableKind {
  return (BeadSpannableList as readonly BeadKind[]).includes(kind);
}

export function getBeadKindTitle(kind: BeadKind): string {
  return getMappedValue(kind, {
    [BeadKind.CIRCLE]: 'Круглий',
    [BeadKind.BUGLE]: 'Склярус',
  });
}
