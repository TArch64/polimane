import { computed, reactive } from 'vue';
import { useEditorStore } from '@editor/stores';
import { reactiveComputed } from '@vueuse/core';
import { BEAD_CENTER, BEAD_SIZE } from '@editor/const';
import {
  type IPoint,
  parseSchemaBeadCoord,
  type SchemaBeadCoord,
  type SchemaBeads,
} from '@/models';
import { getObjectEntries } from '@/helpers';

export interface IBeadsGridItem {
  coord: SchemaBeadCoord;
  offset: IPoint;
  color: string;
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

  const beads = computed(() => (
    getObjectEntries<SchemaBeads>(editorStore.schema.beads)
      .map(([coord, color]): IBeadsGridItem => {
        const { x, y } = resolveBeadOffset(coord);

        return {
          coord,
          color,

          offset: {
            x: x + BEAD_CENTER,
            y: y + BEAD_CENTER,
          },
        };
      })
  ));

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
