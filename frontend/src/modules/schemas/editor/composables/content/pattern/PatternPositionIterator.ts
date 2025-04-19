import type { Canvas } from 'fabric';
import { PatternObject } from '../objects';

interface IPatternPosition {
  object: PatternObject;
  top: number;
  left: number;
}

export class PatternPositionIterator extends Iterator<IPatternPosition, undefined> {
  static readonly CANVAS_PADDING = 10;
  static readonly PATTERN_GAP = 50;

  private readonly availableHorizontalSpace: number;
  private readonly totalHeight: number;
  private index = 0;
  private nextOffsetTop: number;
  readonly [Symbol.iterator] = () => this;

  constructor(
    private readonly canvas: Canvas,
    private readonly objects: PatternObject[],
  ) {
    super();
    this.availableHorizontalSpace = canvas.width - PatternPositionIterator.CANVAS_PADDING * 2;
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
    const freeSpace = this.canvas.height - this.totalHeight - PatternPositionIterator.CANVAS_PADDING * 2;
    return Math.max(freeSpace / 2, PatternPositionIterator.CANVAS_PADDING);
  }

  next(): IteratorResult<IPatternPosition> {
    if (!this.object) {
      return { done: true, value: undefined };
    }

    const value: IPatternPosition = {
      object: this.object,
      top: this.nextOffsetTop,
      left: this.offsetLeft,
    };

    this.nextOffsetTop += this.object.height + PatternPositionIterator.PATTERN_GAP;
    this.index++;
    return { done: false, value };
  }

  private get object(): PatternObject | undefined {
    return this.objects[this.index];
  }

  private get offsetLeft(): number {
    const freeSpaceLeft = this.availableHorizontalSpace - this.object!.width;
    return Math.max(freeSpaceLeft, PatternPositionIterator.CANVAS_PADDING);
  }
}
