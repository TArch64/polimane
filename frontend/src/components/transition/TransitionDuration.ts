export interface ITransitionDuration {
  enter: number;
  leave: number;
}

export type TransitionDuration = number | ITransitionDuration;

export function normalizeDuration(duration: TransitionDuration): ITransitionDuration {
  return typeof duration === 'number' ? { enter: duration, leave: duration } : duration;
}
