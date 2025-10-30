import { computed } from 'vue';
import { EditorCursor, EditorCursorTarget } from '@editor/enums';
import { useCanvasStore } from '@editor/stores';
import type { IEditorTool } from './tool';

export function useNavigateTool() {
  const canvasStore = useCanvasStore();
  let lastEvent: MouseEvent | null = null;

  function onMouseMove(event: MouseEvent) {
    if (!lastEvent) {
      lastEvent = event;
      return;
    }

    const deltaX = lastEvent.clientX - event.clientX;
    const deltaY = lastEvent.clientY - event.clientY;
    canvasStore.navigate(deltaX, deltaY);
    lastEvent = event;
  }

  function onMouseUp() {
    lastEvent = null;
    canvasStore.setCursor(EditorCursor.GRAB, EditorCursorTarget.CANVAS);
    removeEventListener('mousemove', onMouseMove);
  }

  function onMouseDown() {
    canvasStore.setCursor(EditorCursor.GRABBING, EditorCursorTarget.CANVAS);
    addEventListener('mouseup', onMouseUp, { once: true });
    addEventListener('mousemove', onMouseMove);
  }

  return computed((): IEditorTool => ({
    level: 'canvas',
    listeners: { mousedown: onMouseDown },
  }));
}
