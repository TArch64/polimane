import { getObjectKeys } from '@/helpers';

type PaddingSide = 'top' | 'right' | 'bottom' | 'left';
type Padding = Record<PaddingSide, number>;
type PaddingInputSide = PaddingSide | 'horizontal' | 'vertical';
export type PaddingInput = Partial<Record<PaddingInputSide, number>> | number;

export interface IPadding extends Padding {
  horizontal: number;
  vertical: number;
}

function normalizePadding(input: PaddingInput): Padding {
  if (typeof input === 'number') {
    return {
      top: input,
      right: input,
      bottom: input,
      left: input,
    };
  }

  return getObjectKeys(input).reduce((padding, side): Padding => {
    if (side === 'horizontal') {
      return { ...padding, left: input.horizontal!, right: input.horizontal! };
    }
    if (side === 'vertical') {
      return { ...padding, top: input.vertical!, bottom: input.vertical! };
    }
    return { ...padding, [side]: input[side] };
  }, {
    top: 0,
    right: 0,
    bottom: 0,
    left: 0,
  });
}

export function resolvePadding(input: PaddingInput): IPadding {
  const padding = normalizePadding(input);

  return {
    ...padding,
    horizontal: padding.left + padding.right,
    vertical: padding.top + padding.bottom,
  };
}
