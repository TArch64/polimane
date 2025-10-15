import { computed, reactive, type Ref, ref } from 'vue';
import { BEAD_SIZE } from '@editor/const';
import { Direction, DirectionList, isNegativeDirection, isVerticalDirection } from '@/enums';
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
  extendTranslation: (dir: Direction, delta: number) => void;
  reset: () => void;
}

export function useSelectionResize(options: ISelectionResizeOptions): ISelectionResize {
  const beadsStore = useBeadsStore();

  const translation = ref(0);
  const direction = ref<Direction | null>(null);
  const isResizing = computed(() => !!translation.value);

  const fullTranslation = computed(() => Object.fromEntries(
    DirectionList.map((dir) => [dir, dir === direction.value ? translation.value : 0]),
  ) as Record<Direction, number>);

  const resizingX = computed(() => options.area.x - fullTranslation.value.left);
  const resizingY = computed(() => options.area.y - fullTranslation.value.top);
  const resizingWidth = computed(() => options.area.width + fullTranslation.value.right + fullTranslation.value.left);
  const resizingHeight = computed(() => options.area.height + fullTranslation.value.bottom + fullTranslation.value.top);

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

  function extendTranslation(dir: Direction, delta: number) {
    translation.value += delta;

    if (translation.value <= 0) {
      translation.value = 0;
      return;
    }

    if (!direction.value) {
      direction.value = dir;
      const selected = options.selected.value!;

      const from = parseSchemaBeadCoord(selected.from);
      const to = parseSchemaBeadCoord(selected.to);
      const selectedBeads = beadsStore.getInArea(from, to);

      sequenceAxis.value = isVerticalDirection(dir) ? 'y' : 'x';
      capturedSequence.value = buildSequence(selectedBeads, dir);
    }

    const offset = Math.floor(translation.value / BEAD_SIZE);

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
      translation.value -= BEAD_SIZE;
      extendArea(dir, BEAD_SIZE);

      if (sequenceIndex.value === capturedSequence.value.length) {
        sequenceIndex.value = 0;
        sequenceOffset.value += capturedSequence.value.length;
        shift = index;
      }
    }
  }

  function reset(): void {
    translation.value = 0;
    direction.value = null;
    capturedSequence.value = [];
    sequenceAxis.value = null;
    sequenceIndex.value = 0;
    sequenceOffset.value = 0;
  }

  return reactive({
    isResizing,
    translation: fullTranslation,
    extendTranslation,
    direction,
    x: resizingX,
    y: resizingY,
    width: resizingWidth,
    height: resizingHeight,
    reset,
  });
}
