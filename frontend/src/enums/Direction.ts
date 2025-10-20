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

export const VerticalDirectionList = [
  Direction.TOP,
  Direction.BOTTOM,
] as const;

export type VerticalDirection = typeof VerticalDirectionList[number];

export function isVerticalDirection(direction: Direction): direction is VerticalDirection {
  return direction === Direction.TOP || direction === Direction.BOTTOM;
}

export const HorizontalDirectionList = [
  Direction.LEFT,
  Direction.RIGHT,
] as const;

export type HorizontalDirection = typeof HorizontalDirectionList[number];

export function isHorizontalDirection(direction: Direction): direction is HorizontalDirection {
  return direction === Direction.LEFT || direction === Direction.RIGHT;
}

export function isNegativeDirection(direction: Direction): boolean {
  return direction === Direction.TOP || direction === Direction.LEFT;
}
