import type { Ref } from 'vue';
import { util } from 'fabric';
import type { ISchemaPattern } from '@/models';
import { EditorObjectType } from '../../../enums';
import { injectCanvas } from '../../useCanvas';
import { useObjectRenderer } from '../useObjectRenderer';
import { PatternObject } from '../objects';
import { PatternPositionIterator } from './PatternPositionIterator';

export function usePatternRenderer(patterns: Ref<ISchemaPattern[]>) {
  const canvas = injectCanvas();

  useObjectRenderer({
    type: EditorObjectType.PATTERN,
    items: patterns,

    createObject(pattern: ISchemaPattern): PatternObject {
      return new PatternObject(pattern);
    },

    updateObject(item, object) {
      object.update(item);
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
