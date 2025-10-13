import { computed, ref } from 'vue';
import { reactiveComputed } from '@vueuse/core';
import { type INodeRect, NodeRect } from '@/models';

export interface ISelectionArea extends INodeRect {
  setPoint: (x: number, y: number) => void;
  extend: (x: number, y: number) => void;
  reset: () => void;
}

export function useSelectionArea(): ISelectionArea {
  const raw = ref<NodeRect>(NodeRect.BLANK.clone());

  function setPoint(x: number, y: number) {
    raw.value = raw.value.with({ x, y });
  }

  function reset() {
    raw.value = NodeRect.BLANK.clone();
  }

  function extend(x: number, y: number) {
    raw.value.width += x;
    raw.value.height += y;
  }

  const resolved = computed(() => ({
    x: raw.value.width < 0 ? raw.value.x + raw.value.width : raw.value.x,
    y: raw.value.height < 0 ? raw.value.y + raw.value.height : raw.value.y,
    width: Math.abs(raw.value.width),
    height: Math.abs(raw.value.height),
  }));

  return reactiveComputed(() => ({
    ...resolved.value,
    setPoint,
    extend,
    reset,
  }));
}
