export interface ICounter {
  isReached: boolean;
  isOverflowed: boolean;
  current: number;
  max?: number;
  willOverlow: (value: number) => boolean;
}
