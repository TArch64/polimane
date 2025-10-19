import { computed, reactive } from 'vue';
import { useEditorStore } from '@editor/stores';
import { reactiveComputed } from '@vueuse/core';
import { BEAD_CENTER, BEAD_SIZE } from '@editor/const';
import {
  type IPoint,
  isRefBead,
  parseSchemaBeadCoord,
  type SchemaBead,
  type SchemaBeadCoord,
} from '@/models';
import { getObjectEntries } from '@/helpers';
import { type BeadContentKind } from '@/enums';

export interface IBeadsGridItem {
  coord: SchemaBeadCoord;
  offset: IPoint;
  bead: SchemaBead<BeadContentKind>;
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
  resolveBeadOffset: (coord: SchemaBeadCoord) => IPoint;
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

  function resolveBeadOffset(coord: SchemaBeadCoord): IPoint {
    const { x, y } = parseSchemaBeadCoord(coord);
    const offsetX = initialOffsetX + (x * BEAD_SIZE);
    const offsetY = initialOffsetY + (y * BEAD_SIZE);
    return { x: offsetX, y: offsetY };
  }

  const beads = computed(() => {
    const items: IBeadsGridItem[] = [];
    const beads = editorStore.schema.beads;

    for (const [coord, bead] of getObjectEntries(beads)) {
      if (isRefBead(bead)) {
        continue;
      }

      const { x, y } = resolveBeadOffset(coord);

      items.push({
        coord,
        bead,

        offset: {
          x: x + BEAD_CENTER,
          y: y + BEAD_CENTER,
        },
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
