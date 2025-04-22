import type { ObjectParent } from '@/modules/schemas/editor/composables';
import { type IObjectPosition, PositionIterator } from '../PositionIterator';
import { PatternObject } from './PatternObject';

export class PatternPositionIterator extends PositionIterator<PatternObject> {
  static readonly CANVAS_PADDING = 10;
  static readonly PATTERN_GAP = 50;

  private readonly availableHorizontalSpace: number;
  private nextOffsetTop: number;

  constructor(parent: ObjectParent, objects: PatternObject[]) {
    super(parent, objects);
    this.availableHorizontalSpace = parent.width - PatternPositionIterator.CANVAS_PADDING * 2;
    this.nextOffsetTop = this.calcInitialNextOffsetTop();
  }

  private calcInitialNextOffsetTop(): number {
    const freeSpace = this.parent.height - this.totalHeight - PatternPositionIterator.CANVAS_PADDING * 2;
    return Math.max(freeSpace / 2, PatternPositionIterator.CANVAS_PADDING);
  }

  private get totalHeight(): number {
    return this.calcListSize('height', PatternPositionIterator.PATTERN_GAP);
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
    const freeSpaceLeft = (this.availableHorizontalSpace - this.object!.width) / 2;
    return Math.max(freeSpaceLeft, PatternPositionIterator.CANVAS_PADDING);
  }
}
