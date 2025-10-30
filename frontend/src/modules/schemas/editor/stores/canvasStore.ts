import { defineStore } from 'pinia';
import { reactive, ref } from 'vue';
import { clamp } from '@vueuse/core';
import { EditorCursor, EditorCursorTarget } from '@editor/enums';
import type { IPoint } from '@/models';

export const MIN_SCALE = 0.5;
export const MAX_SCALE = 10;

export const useCanvasStore = defineStore('schemas/editor/canvas', () => {
  const scale = ref(1);
  const translation = reactive<IPoint>({ x: 0, y: 0 });
  const cursor = ref<EditorCursor>(EditorCursor.CROSSHAIR);
  const cursorTarget = ref<EditorCursorTarget>(EditorCursorTarget.CONTENT);

  function setScale(value: number): number {
    scale.value = clamp(value, MIN_SCALE, MAX_SCALE);
    return scale.value;
  }

  function navigate(deltaX: number, deltaY: number): void {
    translation.x += deltaX / scale.value;
    translation.y += deltaY / scale.value;
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
    setScale,
    translation,
    navigate,
    cursor,
    cursorTarget,
    setCursor,
  };
});
