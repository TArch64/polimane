export interface ICounter {
  isReached: boolean;
  current: number;
  max?: number;
  willOverlow: (value: number) => boolean;
}
