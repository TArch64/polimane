import { computed, reactive, type Ref, ref } from 'vue';
import { BEAD_SIZE } from '@editor/const';
import {
  BeadKind,
  Direction,
  DirectionList,
  isNegativeDirection,
  isVerticalDirection,
} from '@/enums';
import {
  type BeadCoord,
  type INodeRect,
  isRefBead,
  isSpannableBead,
  parseBeadCoord,
  type SchemaBeads,
  serializeBeadCoord,
  serializeBeadPoint,
} from '@/models';
import { getObjectEntries, getObjectKeys, type ObjectEntries } from '@/helpers';
import { useSchemaBeadsCounter } from '@/composables/subscription';
import { useBeadsStore } from '../beadsStore';
import type { IBeadSelection } from '../selectionStore';
import { useCanvasStore } from '../canvasStore';
import { useEditorStore } from '../editorStore';
import type { ISelectionArea } from './useSelectionArea';
import { useBeadFactory } from './useBeadFactory';

export interface ISelectionResizeOptions {
  area: ISelectionArea;
  selected: Ref<IBeadSelection | null>;
}

export interface ISelectionResize extends INodeRect {
  isResizing: boolean;
  direction: Direction | null;
  translation: Record<Direction, number>;
  extendTranslation: (dir: Direction, delta: number) => void;
  cleanup: () => void;
  reset: () => void;
}

export function useSelectionResize(options: ISelectionResizeOptions): ISelectionResize {
  const editorStore = useEditorStore();
  const beadsStore = useBeadsStore();
  const canvasStore = useCanvasStore();

  const beadFactory = useBeadFactory();
  const limit = useSchemaBeadsCounter(() => editorStore.schema);

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

  function getAxisCoord(coord: BeadCoord): number {
    return parseBeadCoord(coord)[sequenceAxis.value!];
  }

  function buildSequence(beads: SchemaBeads, direction: Direction): SchemaBeads[] {
    const coords = getObjectKeys(beads).map(getAxisCoord);

    const order = isNegativeDirection(direction) ? -1 : 1;
    const mainCoords = Array.from(new Set(coords)).sort((a, b) => (a - b) * order);
    const sequence = mainCoords.map((): SchemaBeads => ({}));

    for (const [coord, bead] of getObjectEntries(beads)) {
      if (isRefBead(bead)) {
        continue;
      }

      const axisIndex = mainCoords.indexOf(getAxisCoord(coord));
      sequence[axisIndex]![coord] = bead;
    }

    return sequence;
  }

  function renderTemplateCoord(templateCoord: BeadCoord): BeadCoord {
    const point = parseBeadCoord(templateCoord);
    const modifier = isNegativeDirection(direction.value!) ? -1 : 1;
    point[sequenceAxis.value!] += (sequenceOffset.value + capturedSequence.value.length) * modifier;
    return serializeBeadPoint(point);
  }

  function renderTemplate(template: SchemaBeads) {
    return getObjectEntries(template).flatMap(([templateCoord, bead]) => {
      const coord = renderTemplateCoord(templateCoord);
      const entries = [[coord, bead]];

      if (!isSpannableBead(bead)) {
        return entries;
      }

      const spanRefs = beadsStore.getSpanBeads<BeadKind.REF>(templateCoord, bead);

      for (const templateRefCoord of getObjectKeys(spanRefs)) {
        entries.push([
          renderTemplateCoord(templateRefCoord),
          beadFactory.createRef(coord),
        ]);
      }

      return entries;
    }) as unknown as ObjectEntries<SchemaBeads>;
  }

  function extendArea(value: number) {
    const isVertical = isVerticalDirection(direction.value!);
    const x = isVertical ? 0 : value;
    const y = isVertical ? value : 0;

    if (isNegativeDirection(direction.value!)) {
      options.area.shiftPoint(-x, -y);
    }

    options.area.extend(x, y);
  }

  function extendSelected(newBeads: ObjectEntries<SchemaBeads>): void {
    const coords = [
      options.selected.value!.from,
      options.selected.value!.to,
      ...newBeads.map(([coord]) => coord),
    ];
    const parsed = coords.map(parseBeadCoord);
    const xs = parsed.map((c) => c.x);
    const ys = parsed.map((c) => c.y);

    options.selected.value = {
      from: serializeBeadCoord(Math.min(...xs), Math.min(...ys)),
      to: serializeBeadCoord(Math.max(...xs), Math.max(...ys)),
    };
  }

  function reset(): void {
    translation.value = 0;
    direction.value = null;
    capturedSequence.value = [];
    sequenceAxis.value = null;
    sequenceIndex.value = 0;
    sequenceOffset.value = 0;
  }

  function extendTranslation(dir: Direction, delta: number) {
    translation.value += delta / canvasStore.scale;

    if (translation.value <= 0) {
      translation.value = 0;
      return;
    }

    if (dir !== direction.value) {
      if (direction.value) {
        reset();
      }

      direction.value = dir;
      const selected = options.selected.value!;

      const from = parseBeadCoord(selected.from);
      const to = parseBeadCoord(selected.to);
      const selectedBeads = beadsStore.getInArea(from, to);

      sequenceAxis.value = isVerticalDirection(dir) ? 'y' : 'x';
      capturedSequence.value = buildSequence(selectedBeads, dir);
    }

    const offset = Math.floor(translation.value / BEAD_SIZE);

    if (offset === 0) {
      return;
    }

    for (let index = 0; index < offset; index++) {
      const templateBeads = capturedSequence.value[sequenceIndex.value]!;
      const newBeads = renderTemplate(templateBeads);

      beadsStore.paintMany(Object.fromEntries(newBeads));
      sequenceIndex.value++;
      translation.value -= BEAD_SIZE;
      extendArea(BEAD_SIZE);
      extendSelected(newBeads);

      if (sequenceIndex.value === capturedSequence.value.length) {
        sequenceIndex.value = 0;
        sequenceOffset.value += capturedSequence.value.length;
      }
    }
  }

  function cleanup() {
    translation.value = 0;
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
    cleanup,
    reset,
  });
}
