import { getMappedValue } from '@/helpers';

export enum BeadKind {
  CIRCLE = 'circle',
  BUGLE = 'bugle',
  REF = 'ref',
}

export const BeadContentList = [BeadKind.CIRCLE, BeadKind.BUGLE] as const;
export type BeadContentKind = typeof BeadContentList[number];

export const BeadSpannableList = [BeadKind.BUGLE] as const;
export type BeadSpannableKind = typeof BeadSpannableList[number];

export function isBeadSpannableKind(kind: BeadKind): kind is BeadSpannableKind {
  return (BeadSpannableList as readonly BeadKind[]).includes(kind);
}

export function getBeadKindTitle(kind: BeadContentKind): string {
  return getMappedValue(kind, {
    [BeadKind.CIRCLE]: 'Круглий',
    [BeadKind.BUGLE]: 'Склярус',
  });
}
