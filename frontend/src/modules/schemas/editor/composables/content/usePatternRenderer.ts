import type { Ref } from 'vue';
import { type Canvas, util } from 'fabric';
import type { ISchemaPattern } from '@/models';
import { EditorObjectType } from '../../enums';
import { injectCanvas } from '../useCanvas';
import { useObjectRenderer } from './useObjectRenderer';
import { PatternObject } from './objects';

const CANVAS_PADDING = 10;
const PATTERN_GAP = 50;

interface IPatternPosition {
  object: PatternObject;
  top: number;
  left: number;
}

class PatternPositionIterator extends Iterator<IPatternPosition, undefined> {
  private readonly freeSpaceLeft: number;
  private readonly totalHeight: number;
  private index: number;
  private nextOffsetTop: number;
  readonly [Symbol.iterator] = () => this;

  constructor(
    canvas: Canvas,
    private readonly objects: PatternObject[],
  ) {
    super();
    this.freeSpaceLeft = canvas.width - CANVAS_PADDING * 2;

    this.totalHeight = objects.reduce((acc, object, index) => {
      return acc + object.height + (index < objects.length ? PATTERN_GAP : 0);
    }, 0);

    this.index = 0;
    this.nextOffsetTop = Math.max((canvas.height - this.totalHeight - CANVAS_PADDING * 2) / 2, CANVAS_PADDING);
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

    this.nextOffsetTop += this.object.height + PATTERN_GAP;
    this.index++;
    return { done: false, value };
  }

  private get object(): PatternObject | undefined {
    return this.objects[this.index];
  }

  private get offsetLeft(): number {
    return Math.max((this.freeSpaceLeft - this.object!.width) / 2, CANVAS_PADDING);
  }
}

export function usePatternRenderer(patterns: Ref<ISchemaPattern[]>) {
  const canvas = injectCanvas();

  useObjectRenderer({
    type: EditorObjectType.PATTERN,
    items: patterns,

    createObject(pattern: ISchemaPattern): PatternObject {
      return new PatternObject(pattern);
    },

    updatePositions(objects: PatternObject[]): void {
      for (const { object, top, left } of new PatternPositionIterator(canvas, objects)) {
        if (object.top !== top) {
          if (object.top > 0) {
            object.animate({ top }, {
              duration: 150,
              onChange: () => canvas.requestRenderAll(),
              easing: util.ease.easeOutQuad,
            });
          } else {
            object.setY(top);
          }
        }

        if (object.left !== left) {
          object.setX(left);
        }
      }
    },
  });
}
