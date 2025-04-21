import type { ObjectParent } from '@/modules/schemas/editor/composables';
import { type IObjectPosition, PositionIterator } from '../PositionIterator';
import { PatternObject } from './PatternObject';

export class PatternPositionIterator extends PositionIterator<PatternObject> {
  static readonly CANVAS_PADDING = 10;
  static readonly PATTERN_GAP = 50;

  private readonly availableHorizontalSpace: number;
  private readonly totalHeight: number;
  private nextOffsetTop: number;

  constructor(parent: ObjectParent, objects: PatternObject[]) {
    super(parent, objects);
    this.availableHorizontalSpace = parent.width - PatternPositionIterator.CANVAS_PADDING * 2;
    this.totalHeight = this.calcTotalHeight();
    this.nextOffsetTop = this.calcInitialNextOffsetTop();
  }

  private calcTotalHeight(): number {
    return this.objects.reduce((acc, object, index) => {
      const gap = index < this.objects.length ? PatternPositionIterator.PATTERN_GAP : 0;
      return acc + object.height + gap;
    }, 0);
  }

  private calcInitialNextOffsetTop(): number {
    const freeSpace = this.parent.height - this.totalHeight - PatternPositionIterator.CANVAS_PADDING * 2;
    return Math.max(freeSpace / 2, PatternPositionIterator.CANVAS_PADDING);
  }

  protected iteration(): IObjectPosition<PatternObject> {
    const value: IObjectPosition<PatternObject> = {
      object: this.object!,
      top: this.nextOffsetTop,
      left: this.offsetLeft,
    };

    this.nextOffsetTop += this.object!.height + PatternPositionIterator.PATTERN_GAP;
    return value;
  }

  private get offsetLeft(): number {
    const freeSpaceLeft = this.availableHorizontalSpace - this.object!.width;
    return Math.max(freeSpaceLeft, PatternPositionIterator.CANVAS_PADDING);
  }
}
