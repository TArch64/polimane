import { defineStore } from 'pinia';
import { reactive, ref } from 'vue';
import { clamp } from '@vueuse/core';
import { EditorCursor, EditorCursorTarget } from '@editor/enums';
import type { IPoint } from '@/models';

export const MIN_SCALE = 0.5;
export const MAX_SCALE = 10;
export const ZOOM_IN_STEP = -50;
export const ZOOM_OUT_STEP = 25;

export const useCanvasStore = defineStore('schemas/editor/canvas', () => {
  const scale = ref(1);
  const translation = reactive<IPoint>({ x: 0, y: 0 });
  const cursor = ref<EditorCursor>(EditorCursor.CROSSHAIR);
  const cursorTarget = ref<EditorCursorTarget>(EditorCursorTarget.CONTENT);

  function navigate(deltaX: number, deltaY: number): void {
    translation.x += deltaX / scale.value;
    translation.y += deltaY / scale.value;
  }

  function zoom(pointX: number, pointY: number, deltaY: number): void {
    const mousePointToX = (pointX / scale.value) + translation.x;
    const mousePointToY = (pointY / scale.value) + translation.y;

    const scaleFactor = 1 - deltaY * 0.01;
    scale.value = clamp(scale.value * scaleFactor, MIN_SCALE, MAX_SCALE);

    translation.x = mousePointToX - (pointX / scale.value);
    translation.y = mousePointToY - (pointY / scale.value);
  }

  function setCursor(
    value: EditorCursor,
    target: EditorCursorTarget = EditorCursorTarget.CONTENT,
  ) {
    cursor.value = value;
    cursorTarget.value = target;
  }

  return {
    scale,
    zoom,
    translation,
    navigate,
    cursor,
    cursorTarget,
    setCursor,
  };
});
