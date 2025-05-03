import { getMappedValue } from '@/helpers';

export const enum PatternType {
  SQUARE = 'square',
  DIAMOND = 'diamond',
}

export const PatternKindValues = [
  PatternType.SQUARE,
  PatternType.DIAMOND,
] as const;

export function getPatternTitle(value: PatternType): string {
  return getMappedValue(value, {
    [PatternType.SQUARE]: 'Квадратна Сітка',
    [PatternType.DIAMOND]: 'Ромбова Сітка',
  });
}
