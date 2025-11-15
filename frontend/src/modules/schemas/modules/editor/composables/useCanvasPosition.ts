import { type MaybeRefOrGetter, reactive, watch } from 'vue';
import { useElementBounding } from '@vueuse/core';
import { useCanvasStore } from '@editor/stores';
import type { INodeRect } from '@/models';

export function useCanvasPosition(reference: MaybeRefOrGetter<SVGElement>): INodeRect {
  const canvasStore = useCanvasStore();

  const { x, y, width, height, update } = useElementBounding(reference);

  watch([
    () => canvasStore.scale,
    () => canvasStore.translation,
  ], update, { deep: true });

  return reactive({ x, y, width, height });
}
