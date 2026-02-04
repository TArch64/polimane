export interface ICounter {
  isReached: boolean;
  isOverflowed: boolean;
  current: number;
  max?: number;
  willOverflow: (value: number) => boolean;
  overflowed: number;
}
