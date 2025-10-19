import { computed, reactive } from 'vue';
import { useEditorStore } from '@editor/stores';
import { reactiveComputed } from '@vueuse/core';
import { BEAD_BUGLE_PADDING, BEAD_CIRCLE_CENTER, BEAD_SIZE } from '@editor/const';
import {
  type BeadCoord,
  type INodeRect,
  type IPoint,
  isRefBead,
  parseBeadCoord,
  type SchemaContentBead,
} from '@/models';
import { getObjectEntries } from '@/helpers';
import { type BeadContentKind, BeadKind } from '@/enums';

export interface IBeadsGridCircle {
  center: IPoint;
}

export interface IBeadsGridBugle extends INodeRect {
}

type BeadPrecomutedDataMap = {
  [BeadKind.CIRCLE]: IBeadsGridCircle;
  [BeadKind.BUGLE]: IBeadsGridBugle;
};

interface IComputeBeadDataOptions {
  coord: IPoint;
  bead: SchemaContentBead;
}

type ComputeBeadData<K extends BeadContentKind> = (options: IComputeBeadDataOptions) => BeadPrecomutedDataMap[K];
type BeadDataComputers = { [K in BeadContentKind]: ComputeBeadData<K> };

export interface IBeadsGridItem {
  coord: BeadCoord;
  precomputed: BeadPrecomutedDataMap[BeadContentKind];
  bead: SchemaContentBead;
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
      x: options.coord.x + (span.x),
      y: options.coord.y + (span.y),
    };

    const startCoordX = Math.min(options.coord.x, spanCoord.x);
    const startCoordY = Math.min(options.coord.y, spanCoord.y);
    const coordWidth = Math.abs(span.x) + 1;
    const coordHeight = Math.abs(span.y) + 1;

    const offset = resolveBeadOffset({ x: startCoordX, y: startCoordY });
    const width = coordWidth * BEAD_SIZE - BEAD_BUGLE_PADDING * 2;
    const height = coordHeight * BEAD_SIZE - BEAD_BUGLE_PADDING * 2;

    return {
      x: offset.x + BEAD_BUGLE_PADDING,
      y: offset.y + BEAD_BUGLE_PADDING,
      width,
      height,
    };
  };

  const computers: BeadDataComputers = {
    [BeadKind.CIRCLE]: computeBeadCircle,
    [BeadKind.BUGLE]: computeBeadBugle,
  };

  const beads = computed(() => {
    const items: IBeadsGridItem[] = [];
    const beads = editorStore.schema.beads;

    for (const [coord, bead_] of getObjectEntries(beads)) {
      if (isRefBead(bead_)) {
        continue;
      }

      const bead = bead_ as SchemaContentBead;
      const parsedCoord = parseBeadCoord(coord);

      items.push({
        coord,
        bead,

        precomputed: computers[bead.kind as BeadContentKind]({
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
