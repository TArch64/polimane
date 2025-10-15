import { computed, reactive, ref } from 'vue';
import { type INodeRect } from '@/models';

export interface ISelectionArea extends INodeRect {
  setPoint: (x: number, y: number) => void;
  shiftPoint: (dx: number, dy: number) => void;
  extend: (dx: number, dy: number) => void;
  reset: () => void;
}

export function useSelectionArea(): ISelectionArea {
  const x = ref(0);
  const y = ref(0);
  const width = ref(0);
  const height = ref(0);

  function reset() {
    x.value = 0;
    y.value = 0;
    width.value = 0;
    height.value = 0;
  }

  function setPoint(x_: number, y_: number) {
    reset();
    x.value = x_;
    y.value = y_;
  }

  function shiftPoint(dx: number, dy: number) {
    x.value += dx;
    y.value += dy;
  }

  function extend(dx: number, dy: number) {
    width.value += dx;
    height.value += dy;
  }

  const resolvedX = computed(() => width.value < 0 ? x.value + width.value : x.value);
  const resolvedY = computed(() => height.value < 0 ? y.value + height.value : y.value);
  const resolvedWidth = computed(() => Math.abs(width.value));
  const resolvedHeight = computed(() => Math.abs(height.value));

  return reactive({
    x: resolvedX,
    y: resolvedY,
    width: resolvedWidth,
    height: resolvedHeight,
    setPoint,
    shiftPoint,
    extend,
    reset,
  });
}
