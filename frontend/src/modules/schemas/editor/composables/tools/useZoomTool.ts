import { computed } from 'vue';
import { EditorCursor, EditorCursorTarget } from '@editor/enums';
import { useCanvasStore, ZOOM_IN_STEP, ZOOM_OUT_STEP } from '@editor/stores';
import type { IEditorTool } from './tool';

export const useZoomTool = () => {
  const canvasStore = useCanvasStore();
  let isZoomOut = false;

  function onClick(event: PointerEvent) {
    const deltaY = isZoomOut ? ZOOM_OUT_STEP : ZOOM_IN_STEP;
    canvasStore.zoom(event.clientX, event.clientY, deltaY);
  }

  function onKeyDown(event: KeyboardEvent) {
    if (event.shiftKey) {
      isZoomOut = true;
      canvasStore.setCursor(EditorCursor.ZOOM_OUT, EditorCursorTarget.CANVAS);
    }
  }

  function onKeyUp() {
    if (isZoomOut) {
      isZoomOut = false;
      canvasStore.setCursor(EditorCursor.ZOOM_IN, EditorCursorTarget.CANVAS);
    }
  }

  function onActivated(abortController: AbortController) {
    addEventListener('keydown', onKeyDown, { signal: abortController.signal });
    addEventListener('keyup', onKeyUp, { signal: abortController.signal });
  }

  return computed((): IEditorTool => ({
    level: 'canvas',
    onActivated,
    listeners: { click: onClick },
  }));
};
