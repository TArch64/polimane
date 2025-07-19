import type { Ref } from 'vue';
import Konva from 'konva';

export function useCanvasStage() {
  return window.__KONVA_STAGE_REF__ as Ref<Konva.Stage>;
}
