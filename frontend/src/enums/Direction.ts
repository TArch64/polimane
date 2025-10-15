export enum Direction {
  TOP = 'top',
  LEFT = 'left',
  RIGHT = 'right',
  BOTTOM = 'bottom',
}

export const DirectionList = [
  Direction.TOP,
  Direction.LEFT,
  Direction.RIGHT,
  Direction.BOTTOM,
] as const;

export function isVerticalDirection(direction: Direction): boolean {
  return direction === Direction.TOP || direction === Direction.BOTTOM;
}

export function isNegativeDirection(direction: Direction): boolean {
  return direction === Direction.TOP || direction === Direction.LEFT;
}
