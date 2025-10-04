import { computed, type MaybeRefOrGetter, reactive, toValue } from 'vue';
import { type ISchema, parseSchemaBeadCoord, type SchemaBeadCoord } from '@/models';

export const BEAD_SIZE = 12;
export const BEAD_CENTER = BEAD_SIZE / 2;
export const BEAD_RADIUS = BEAD_CENTER - 1;

export type BeadOffset = [x: number, y: number];

export interface IBeadsGridItem {
  coord: SchemaBeadCoord;
  offset: BeadOffset;
  color: string;
}

export interface IBeadsGrid {
  beads: IBeadsGridItem[];
  size: {
    minX: number;
    minY: number;
    width: number;
    height: number;
  };
}

export function useBeadsGrid(schemaRef: MaybeRefOrGetter<ISchema>): IBeadsGrid {
  const schema = computed(() => toValue(schemaRef));
  const left = computed(() => schema.value.size.left);
  const top = computed(() => schema.value.size.top);
  const right = computed(() => schema.value.size.right);
  const bottom = computed(() => schema.value.size.bottom);

  const initialOffsetX = left.value * BEAD_SIZE;
  const initialOffsetY = top.value * BEAD_SIZE;

  const minOffsetX = computed(() => initialOffsetX - (left.value * BEAD_SIZE));
  const minOffsetY = computed(() => initialOffsetY - (top.value * BEAD_SIZE));

  const width = computed(() => (left.value + right.value) * BEAD_SIZE);
  const height = computed(() => (top.value + bottom.value) * BEAD_SIZE);

  const beads = computed(() => (
    Object.entries(schema.value.beads).map(([coord_, color]): IBeadsGridItem => {
      const coord = coord_ as SchemaBeadCoord;
      const [x, y] = parseSchemaBeadCoord(coord);
      const offsetX = initialOffsetX + (x * BEAD_SIZE);
      const offsetY = initialOffsetY + (y * BEAD_SIZE);

      return {
        coord,
        offset: [offsetX, offsetY],
        color,
      };
    })
  ));

  return reactive({
    beads,

    size: {
      minX: minOffsetX,
      minY: minOffsetY,
      width,
      height,
    },
  });
}
