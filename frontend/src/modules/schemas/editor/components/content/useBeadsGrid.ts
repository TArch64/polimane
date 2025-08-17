import { computed, type ComputedRef } from 'vue';
import type { SchemaBeedCoord } from '@/models';
import { useEditorStore } from '../../stores';

export const BEAD_SIZE = 12;

export type BeadOffset = [x: number, y: number];

export interface IBeadsGridItem {
  coord: SchemaBeedCoord;
  offset: BeadOffset;
}

export interface IBeadsGrid {
  sector: 'topLeft' | 'topRight' | 'bottomLeft' | 'bottomRight';
  grid: ComputedRef<IBeadsGridItem[]>;
}

export function useBeadsGrid(): IBeadsGrid[] {
  const editorStore = useEditorStore();
  const left = computed(() => -editorStore.schema.size.left);
  const top = computed(() => -editorStore.schema.size.top);
  const right = computed(() => editorStore.schema.size.right);
  const bottom = computed(() => editorStore.schema.size.bottom);

  function* grid(fromX: number, toX: number, fromY: number, toY: number): Generator<IBeadsGridItem, void, unknown> {
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
