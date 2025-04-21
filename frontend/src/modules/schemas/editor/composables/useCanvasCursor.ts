import type { FabricObject } from 'fabric';
import type { BrowserCursor } from '@/types';
import { injectCanvas } from './useCanvas';

export interface ICanvasCursor {
  change: (cursor: BrowserCursor, affectedObject?: FabricObject) => void;
  changeTemporarily: (cursor: BrowserCursor, timeout: number, affectedObject?: FabricObject) => void;
}

export function useCanvasCursor(): ICanvasCursor {
  let timeoutId: TimeoutId | null = null;
  const canvas = injectCanvas();

  function change(cursor: BrowserCursor, affectedObject?: FabricObject): void {
    canvas.defaultCursor = cursor;
    canvas.setCursor(cursor);

    if (affectedObject) {
      affectedObject.hoverCursor = cursor;
    }

    canvas.requestRenderAll();
  }

  function changeTemporarily(cursor: BrowserCursor, timeout: number, affectedObject?: FabricObject): void {
    change(cursor, affectedObject);

    if (timeoutId) {
      clearTimeout(timeoutId);
    }

    timeoutId = setTimeout(() => {
      change('default');
      timeoutId = null;
    }, timeout);
  }

  return { change, changeTemporarily };
}
