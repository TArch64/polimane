import type { Ref } from 'vue';
import { type FabricObject, util } from 'fabric';
import type { ISchemaPattern } from '@/models';
import { EditorObjectType } from '../../enums';
import { injectCanvas } from '../useCanvas';
import { useObjectRenderer } from './useObjectRenderer';
import { PatternObject } from './objects';

const CANVAS_PADDING = 10;
const PATTERN_GAP = 50;

export function usePatternRenderer(patterns: Ref<ISchemaPattern[]>) {
  const canvas = injectCanvas();

  useObjectRenderer({
    type: EditorObjectType.PATTERN,
    items: patterns,

    createObject(pattern: ISchemaPattern): PatternObject {
      return new PatternObject(pattern);
    },

    updatePositions(objects: FabricObject[]): void {
      const freeSpaceX = canvas.width - CANVAS_PADDING * 2;

      const totalHeight = objects.reduce((acc, object, index) => {
        return acc + object.height + (index < objects.length ? PATTERN_GAP : 0);
      }, 0);

      let nextOffsetY = Math.max((canvas.height - totalHeight - CANVAS_PADDING * 2) / 2, CANVAS_PADDING);

      for (const object of objects) {
        if (object.top !== nextOffsetY) {
          if (object.top > 0) {
            object.animate({ top: nextOffsetY }, {
              duration: 150,
              onChange: () => canvas.requestRenderAll(),
              easing: util.ease.easeOutQuad,
            });
          } else {
            object.setY(nextOffsetY);
          }
        }

        const offsetLeft = Math.max((freeSpaceX - object.width) / 2, CANVAS_PADDING);

        if (object.left !== offsetLeft) {
          object.setX(offsetLeft);
        }

        nextOffsetY += object.height + PATTERN_GAP;
      }
    },
  });
}
