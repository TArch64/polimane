import { computed, reactive, type Ref, ref, watch } from 'vue';
import { BEAD_SIZE } from '@editor/const';
import { watchThrottled } from '@vueuse/core';
import { Direction, isNegativeDirection, isVerticalDirection } from '@/enums';
import {
  type INodeRect,
  parseSchemaBeadCoord,
  type SchemaBeadCoord,
  type SchemaBeads,
  serializeSchemaBeadCoord,
} from '@/models';
import { getObjectEntries, getObjectKeys } from '@/helpers';
import { useBeadsStore } from '../beadsStore';
import type { IBeadSelection } from '../selectionStore';
import type { ISelectionArea } from './useSelectionArea';

export interface ISelectionResizeOptions {
  area: ISelectionArea;
  selected: Ref<IBeadSelection | null>;
}

export interface ISelectionResize extends INodeRect {
  isResizing: boolean;
  direction: Direction | null;
  translation: Record<Direction, number>;
  reset: () => void;
}

export function useSelectionResize(options: ISelectionResizeOptions): ISelectionResize {
  const beadsStore = useBeadsStore();

  const translation = reactive<Record<Direction, number>>({
    [Direction.TOP]: 0,
    [Direction.BOTTOM]: 0,
    [Direction.LEFT]: 0,
    [Direction.RIGHT]: 0,
  });

  const direction = computed(() => {
    return getObjectKeys(translation)
      .find((direction) => !!translation[direction])
      || null;
  });

  const directionTranslation = computed(() => direction.value ? translation[direction.value] : 0);

  const isResizing = computed(() => !!direction.value);
  const resizingX = computed(() => options.area.x - translation.left);
  const resizingY = computed(() => options.area.y - translation.top);
  const resizingWidth = computed(() => options.area.width + translation.right + translation.left);
  const resizingHeight = computed(() => options.area.height + translation.bottom + translation.top);

  const capturedSequence = ref<SchemaBeads[]>([]);
  const sequenceAxis = ref<'x' | 'y' | null>(null);
  const sequenceIndex = ref(0);
  const sequenceOffset = ref(0);

  function getAxisCoord(coord: SchemaBeadCoord): number {
    return parseSchemaBeadCoord(coord)[sequenceAxis.value!];
  }

  function buildSequence(beads: SchemaBeads, direction: Direction): SchemaBeads[] {
    const order = isNegativeDirection(direction) ? -1 : 1;
    const coords = getObjectKeys(beads).map(getAxisCoord);
    const mainCoords = Array.from(new Set(coords)).sort((a, b) => (a - b) * order);
    const sequence = mainCoords.map((): SchemaBeads => ({}));

    for (const [coord, color] of getObjectEntries(beads)) {
      const axisIndex = mainCoords.indexOf(getAxisCoord(coord));
      sequence[axisIndex]![coord] = color;
    }

    return sequence;
  }

  watch(direction, (direction) => {
    const selected = options.selected.value;

    if (!direction || !selected) {
      return;
    }

    const from = parseSchemaBeadCoord(selected.from);
    const to = parseSchemaBeadCoord(selected.to);
    const selectedBeads = beadsStore.getInArea(from, to);

    sequenceAxis.value = isVerticalDirection(direction) ? 'y' : 'x';
    capturedSequence.value = buildSequence(selectedBeads, direction);
  });

  const sequencePendingOffset = computed(() => {
    return Math.floor(directionTranslation.value / BEAD_SIZE);
  });

  function renderTemplate(template: SchemaBeads) {
    return getObjectEntries<SchemaBeads>(template).map(([templateCoord, bead]) => {
      const coord = parseSchemaBeadCoord(templateCoord);
      coord[sequenceAxis.value!] += sequenceOffset.value + capturedSequence.value.length;
      return [serializeSchemaBeadCoord(coord.x, coord.y), bead] as const;
    });
  }

  function extendArea(direction: Direction, value: number) {
    const mainAxis = isNegativeDirection(direction) ? -value : value;
    const x = isVerticalDirection(direction) ? 0 : mainAxis;
    const y = isVerticalDirection(direction) ? mainAxis : 0;
    options.area.extend(x, y);
  }

  watchThrottled(sequencePendingOffset, (offset) => {
    if (offset === 0) {
      return;
    }

    let shift = 0;

    for (let index = 0; index < offset; index++) {
      const templateIndex = sequenceIndex.value + index - shift;
      const templateBeads = capturedSequence.value[templateIndex]!;
      const newBeads = renderTemplate(templateBeads);

      for (const [coord, bead] of newBeads) {
        beadsStore.paint(coord, bead);
      }

      sequenceIndex.value++;
      translation[direction.value!] -= BEAD_SIZE;
      extendArea(direction.value!, BEAD_SIZE);

      if (sequenceIndex.value === capturedSequence.value.length) {
        sequenceIndex.value = 0;
        sequenceOffset.value += capturedSequence.value.length;
        shift = index;
      }
    }
  }, {
    throttle: 10,
    trailing: true,
  });

  function reset(): void {
    const entries = getObjectKeys(translation).map((direction) => [direction, 0]);
    Object.assign(translation, Object.fromEntries(entries));
    capturedSequence.value = [];
    sequenceAxis.value = null;
    sequenceIndex.value = 0;
    sequenceOffset.value = 0;
  }

  return reactive({
    isResizing,
    translation,
    direction,
    x: resizingX,
    y: resizingY,
    width: resizingWidth,
    height: resizingHeight,
    reset,
  });
}
