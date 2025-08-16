import { computed, type ComputedRef } from 'vue';
import type { SchemaBeedCoordinate } from '@/models';
import { useEditorStore } from '../../stores';

export const BEAD_SIZE = 16;

export type BeadOffset = [x: number, y: number];

export interface IBeadsGridItem {
  coord: SchemaBeedCoordinate;
  offset: BeadOffset;
}

export type BeadGridGenerator = Generator<IBeadsGridItem, void, unknown>;
export type BeadGridSector = 'topLeft' | 'topRight' | 'bottomLeft' | 'bottomRight';

export interface IBeadsGrid {
  sector: BeadGridSector;
  grid: ComputedRef<BeadGridGenerator>;
}

export function useBeadsGrid(): IBeadsGrid[] {
  const editorStore = useEditorStore();
  const left = computed(() => -editorStore.schema.size.left);
  const top = computed(() => -editorStore.schema.size.top);
  const right = computed(() => editorStore.schema.size.right);
  const bottom = computed(() => editorStore.schema.size.bottom);

  function* grid(fromX: number, toX: number, fromY: number, toY: number): BeadGridGenerator {
    const initialOffsetX = (right.value - left.value) * BEAD_SIZE;
    const initialOffsetY = (bottom.value - top.value) * BEAD_SIZE;

    for (let x = fromX; x <= toX; x++) {
      for (let y = fromY; y <= toY; y++) {
        const absoluteX = left.value + x;
        const absoluteY = top.value + y;
        const offsetX = initialOffsetX + (absoluteX * BEAD_SIZE);
        const offsetY = initialOffsetY + (absoluteY * BEAD_SIZE);

        yield {
          coord: `${x}:${y}`,
          offset: [offsetX, offsetY],
        };
      }
    }
  }

  return [
    {
      sector: 'topLeft',
      grid: computed(() => grid(left.value, -1, top.value, -1)),
    },
    {
      sector: 'topRight',
      grid: computed(() => grid(0, right.value, top.value, -1)),
    },
    {
      sector: 'bottomLeft',
      grid: computed(() => grid(left.value, -1, 0, bottom.value)),
    },
    {
      sector: 'bottomRight',
      grid: computed(() => grid(0, right.value, 0, bottom.value)),
    },
  ];
}
