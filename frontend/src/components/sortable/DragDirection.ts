import type { Edge } from '@atlaskit/pragmatic-drag-and-drop-hitbox/types';

export type DragDirection = 'horizontal' | 'vertical';

export function getBeforeDirection(direction: DragDirection): Edge {
  return direction === 'horizontal' ? 'left' : 'top';
}

export function getAfterDirection(direction: DragDirection): Edge {
  return direction === 'horizontal' ? 'right' : 'bottom';
}
