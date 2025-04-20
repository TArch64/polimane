import { getMapValue } from '@/helpers';

export const enum PatternKind {
  SQUARE = 'square',
  DIAMOND = 'diamond',
}

export const PatternKindValues = [
  PatternKind.SQUARE,
  PatternKind.DIAMOND,
] as const;

export function getPatternTitle(value: PatternKind): string {
  return getMapValue(value, {
    [PatternKind.SQUARE]: 'Квадратна Сітка',
    [PatternKind.DIAMOND]: 'Ромбова Сітка',
  });
}
