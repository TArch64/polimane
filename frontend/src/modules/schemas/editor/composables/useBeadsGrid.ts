import { computed, reactive } from 'vue';
import { useEditorStore } from '@editor/stores';
import { parseSchemaBeadCoord, type SchemaBeadCoord } from '@/models';

export const BEAD_SIZE = 12;
export const BEAD_CENTER = BEAD_SIZE / 2;
export const BEAD_RADIUS = BEAD_CENTER - 1;

export type BeadOffset = [x: number, y: number];

export interface IBeadsGridItem {
  coord: SchemaBeadCoord;
  offset: BeadOffset;
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
  resolveBeadOffset: (coord: SchemaBeadCoord) => BeadOffset;
}

export function useBeadsGrid(): IBeadsGrid {
  const editorStore = useEditorStore();

  const left = computed(() => editorStore.schema.size.left);
  const top = computed(() => editorStore.schema.size.top);
  const right = computed(() => editorStore.schema.size.right);
  const bottom = computed(() => editorStore.schema.size.bottom);

  const initialOffsetX = left.value * BEAD_SIZE;
  const initialOffsetY = top.value * BEAD_SIZE;

  const minOffsetX = computed(() => initialOffsetX - (left.value * BEAD_SIZE));
  const minOffsetY = computed(() => initialOffsetY - (top.value * BEAD_SIZE));

  const width = computed(() => (left.value + right.value) * BEAD_SIZE);
  const height = computed(() => (top.value + bottom.value) * BEAD_SIZE);

  function resolveBeadOffset(coord: SchemaBeadCoord): BeadOffset {
    const [x, y] = parseSchemaBeadCoord(coord);
    const offsetX = initialOffsetX + (x * BEAD_SIZE);
    const offsetY = initialOffsetY + (y * BEAD_SIZE);
    return [offsetX, offsetY];
  }

  const beads = computed(() => (
    Object.entries(editorStore.schema.beads)
      .map(([coord_, color]): IBeadsGridItem => {
        const coord = coord_ as SchemaBeadCoord;

        return {
          coord,
          color,
          offset: resolveBeadOffset(coord),
        };
      })
  ));

  return reactive({
    beads,
    resolveBeadOffset,

    size: {
      minX: minOffsetX,
      minY: minOffsetY,
      width,
      height,
    },
  });
}
