import { computed, reactive } from 'vue';
import { useEditorStore } from '@editor/stores';
import { reactiveComputed } from '@vueuse/core';
import {
  BEAD_BUGLE_PADDING,
  BEAD_CIRCLE_CENTER,
  BEAD_REF_HITBOX_PADDING,
  BEAD_SIZE,
} from '@editor/const';
import {
  type BeadCoord,
  type INodeRect,
  type IPoint,
  parseBeadCoord,
  type SchemaBead,
} from '@/models';
import { getObjectEntries } from '@/helpers';
import { BeadKind } from '@/enums';

export interface IBeadsGridCircle {
  center: IPoint;
}

export interface IBeadsGridBugle {
  shape: INodeRect;
}

export interface IBeadsGridRef {
  hitbox: IPoint;
}

type BeadPrecomputedDataMap = {
  [BeadKind.CIRCLE]: IBeadsGridCircle;
  [BeadKind.BUGLE]: IBeadsGridBugle;
  [BeadKind.REF]: IBeadsGridRef;
};

interface IComputeBeadDataOptions {
  coord: IPoint;
  bead: SchemaBead;
}

type ComputeBeadData<K extends BeadKind> = (options: IComputeBeadDataOptions) => BeadPrecomputedDataMap[K];
type BeadDataComputers = { [K in BeadKind]: ComputeBeadData<K> };

export interface IBeadsGridItem {
  coord: BeadCoord;
  precomputed: BeadPrecomputedDataMap[BeadKind];
  bead: SchemaBead;
}

export interface IBeadsGridSize {
  minX: number;
  minY: number;
  width: number;
  height: number;
}

export interface IBeadsGrid {
  beads: IBeadsGridItem[];
  size: IBeadsGridSize;
  resolveBeadOffset: (coord: BeadCoord | IPoint) => IPoint;
}

export function useBeadsGrid(): IBeadsGrid {
  const editorStore = useEditorStore();
  const size = reactiveComputed(() => editorStore.schema.size);

  const initialOffsetX = size.left * BEAD_SIZE;
  const initialOffsetY = size.top * BEAD_SIZE;

  const minX = computed(() => initialOffsetX - (size.left * BEAD_SIZE));
  const minY = computed(() => initialOffsetY - (size.top * BEAD_SIZE));

  const width = computed(() => (size.left + size.right) * BEAD_SIZE);
  const height = computed(() => (size.top + size.bottom) * BEAD_SIZE);

  function resolveBeadOffset(coord_: BeadCoord | IPoint): IPoint {
    const coord = typeof coord_ === 'string' ? parseBeadCoord(coord_) : coord_;
    const offsetX = initialOffsetX + (coord.x * BEAD_SIZE);
    const offsetY = initialOffsetY + (coord.y * BEAD_SIZE);
    return { x: offsetX, y: offsetY };
  }

  const computeBeadCircle: ComputeBeadData<BeadKind.CIRCLE> = (options) => {
    const offset = resolveBeadOffset(options.coord);

    return {
      center: {
        x: offset.x + BEAD_CIRCLE_CENTER,
        y: offset.y + BEAD_CIRCLE_CENTER,
      },
    };
  };

  const computeBeadBugle: ComputeBeadData<BeadKind.BUGLE> = (options) => {
    const span = options.bead.bugle!.span;

    const spanCoord: IPoint = {
      x: options.coord.x + span.x,
      y: options.coord.y + span.y,
    };

    const startCoordX = Math.min(options.coord.x, spanCoord.x);
    const startCoordY = Math.min(options.coord.y, spanCoord.y);
    const coordWidth = Math.abs(span.x) + 1;
    const coordHeight = Math.abs(span.y) + 1;

    const offset = resolveBeadOffset({ x: startCoordX, y: startCoordY });
    const width = coordWidth * BEAD_SIZE - BEAD_BUGLE_PADDING * 2;
    const height = coordHeight * BEAD_SIZE - BEAD_BUGLE_PADDING * 2;

    return {
      shape: {
        x: offset.x + BEAD_BUGLE_PADDING,
        y: offset.y + BEAD_BUGLE_PADDING,
        width,
        height,
      },
    };
  };

  const computeBeadRef: ComputeBeadData<BeadKind.REF> = (options) => {
    const offset = resolveBeadOffset(options.coord);

    return {
      hitbox: {
        x: offset.x + BEAD_REF_HITBOX_PADDING,
        y: offset.y + BEAD_REF_HITBOX_PADDING,
      },
    };
  };

  const computers: BeadDataComputers = {
    [BeadKind.CIRCLE]: computeBeadCircle,
    [BeadKind.BUGLE]: computeBeadBugle,
    [BeadKind.REF]: computeBeadRef,
  };

  const beads = computed(() => {
    const items: IBeadsGridItem[] = [];
    const beads = editorStore.schema.beads;

    for (const [coord, bead] of getObjectEntries(beads)) {
      const parsedCoord = parseBeadCoord(coord);

      items.push({
        coord,
        bead,

        precomputed: computers[bead.kind]({
          coord: parsedCoord,
          bead,
        }),
      });
    }

    return items;
  });

  return reactive({
    beads,
    resolveBeadOffset,

    size: {
      minX: minX,
      minY: minY,
      width,
      height,
    },
  });
}
