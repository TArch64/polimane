import { getMappedValue } from '@/helpers';

export enum AccessLevel {
  READ = 1,
  WRITE = 2,
  ADMIN = 3,
}

export const AccessLeveList = [
  AccessLevel.READ,
  AccessLevel.WRITE,
  AccessLevel.ADMIN,
] as const;

export function getAccessLevelTitle(level: AccessLevel): string {
  return getMappedValue(level, {
    [AccessLevel.READ]: 'Перегляд',
    [AccessLevel.WRITE]: 'Редагування',
    [AccessLevel.ADMIN]: 'Повний',
  });
}
