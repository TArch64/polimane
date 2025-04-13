import type { FabricObject } from 'fabric';
import type { BrowserCursor } from '@/types';
import { injectCanvas } from './useCanvas';

export interface ICanvasCursor {
  change: (cursor: BrowserCursor, affectedObject?: FabricObject) => void;
}

export function useCanvasCursor(): ICanvasCursor {
  const canvas = injectCanvas();

  function change(cursor: BrowserCursor, affectedObject?: FabricObject): void {
    canvas.value.defaultCursor = cursor;
    canvas.value.setCursor(cursor);

    if (affectedObject) {
      affectedObject.hoverCursor = cursor;
    }

    canvas.value.requestRenderAll();
  }

  return { change };
}
