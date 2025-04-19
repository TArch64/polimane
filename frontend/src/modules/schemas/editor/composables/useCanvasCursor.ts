import type { FabricObject } from 'fabric';
import type { BrowserCursor } from '@/types';
import { injectCanvas } from './useCanvas';

export interface ICanvasCursor {
  change: (cursor: BrowserCursor, affectedObject?: FabricObject) => void;
}

export function useCanvasCursor(): ICanvasCursor {
  const canvas = injectCanvas();

  function change(cursor: BrowserCursor, affectedObject?: FabricObject): void {
    canvas.defaultCursor = cursor;
    canvas.setCursor(cursor);

    if (affectedObject) {
      affectedObject.hoverCursor = cursor;
    }

    canvas.requestRenderAll();
  }

  return { change };
}
