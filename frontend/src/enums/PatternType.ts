import { getMapValue } from '@/helpers';

export const enum PatternType {
  SQUARE = 'square',
  DIAMOND = 'diamond',
}

export const PatternTypeValues = [
  PatternType.SQUARE,
  PatternType.DIAMOND,
] as const;

export function getPatternTitle(value: PatternType): string {
  return getMapValue(value, {
    [PatternType.SQUARE]: 'Квадратна Сітка',
    [PatternType.DIAMOND]: 'Ромбова Сітка',
  });
}
