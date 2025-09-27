import { computed, type ComputedRef, type MaybeRefOrGetter, toValue } from 'vue';
import { type ISchema, type SchemaBeadCoord, serializeSchemaBeadCoord } from '@/models';

export const BEAD_SIZE = 12;

export type BeadOffset = [x: number, y: number];

export interface IBeadsGridItem {
  coord: SchemaBeadCoord;
  offset: BeadOffset;
}

export interface IBeadsGrid {
  sector: 'topLeft' | 'topRight' | 'bottomLeft' | 'bottomRight';
  grid: ComputedRef<IBeadsGridItem[]>;
}

export interface IBeadsGridOptions {
  filter?: (coord: SchemaBeadCoord) => boolean;
}

export function useBeadsGrid(schemaRef: MaybeRefOrGetter<ISchema>, options: IBeadsGridOptions = {}): IBeadsGrid[] {
  const filter = options.filter ?? (() => true);

  const size = computed(() => toValue(schemaRef).size);
  const left = computed(() => -size.value.left);
  const top = computed(() => -size.value.top);
  const right = computed(() => size.value.right);
  const bottom = computed(() => size.value.bottom);

  const initialOffsetX = -left.value * BEAD_SIZE;
  const initialOffsetY = -top.value * BEAD_SIZE;

  function* grid(fromX: number, toX: number, fromY: number, toY: number): Generator<IBeadsGridItem, void, unknown> {
    for (let x = fromX; x <= toX; x++) {
      for (let y = fromY; y <= toY; y++) {
        const coord = serializeSchemaBeadCoord(x, y);

        if (!filter(coord)) continue;

        const offsetX = initialOffsetX + (x * BEAD_SIZE);
        const offsetY = initialOffsetY + (y * BEAD_SIZE);

        yield { coord, offset: [offsetX, offsetY] };
      }
    }
  }

  return [
    {
      sector: 'topLeft',
      grid: computed(() => Array.from(grid(left.value, 0, top.value, 0))),
    },
    {
      sector: 'topRight',
      grid: computed(() => Array.from(grid(1, right.value, top.value, 0))),
    },
    {
      sector: 'bottomLeft',
      grid: computed(() => Array.from(grid(left.value, 0, 1, bottom.value))),
    },
    {
      sector: 'bottomRight',
      grid: computed(() => Array.from(grid(1, right.value, 1, bottom.value))),
    },
  ];
}
