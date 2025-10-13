import { computed, type Ref, watch } from 'vue';
import {
  type IBeadSelection,
  useBeadsStore,
  useSelectionStore,
  useToolsStore,
} from '@editor/stores';
import { type IPoint, parseSchemaBeadCoord, Point, serializeSchemaBeadCoord } from '@/models';
import { getObjectKeys } from '@/helpers';
import type { IBeadToolsOptions } from './IBeadToolsOptions';
import { type IBeadResolveOptions, useBeadCoord } from './useBeadCoord';

export interface IBeadSelectionListeners {
  mousedown: (event: MouseEvent) => void;
}

export function useBeadSelection(options: IBeadToolsOptions): Ref<IBeadSelectionListeners> {
  const selectionStore = useSelectionStore();
  const toolsStore = useToolsStore();
  const beadsStore = useBeadsStore();

  const beadCoord = useBeadCoord(options);

  function createSelection(from: IPoint, to: IPoint): IBeadSelection | null {
    const selected = getObjectKeys(beadsStore.getInArea(from, to));

    if (!selected.length) {
      return null;
    }

    if (selected.length === 1) {
      const coord = selected[0]!;
      return { from: coord, to: coord };
    }

    const xs: number[] = [];
    const ys: number[] = [];

    for (const coord of selected.map(parseSchemaBeadCoord)) {
      xs.push(coord.x);
      ys.push(coord.y);
    }

    return {
      from: serializeSchemaBeadCoord(Math.min(...xs), Math.min(...ys)),
      to: serializeSchemaBeadCoord(Math.max(...xs), Math.max(...ys)),
    };
  }

  function onMouseMove(event: MouseEvent) {
    selectionStore.area.extend(event.movementX, event.movementY);
  }

  function onMouseup() {
    try {
      removeEventListener('mousemove', onMouseMove);
      beadCoord.clearCache();

      const { x, y, width, height } = selectionStore.area;
      let point = new Point({ x, y });
      const resolveOptions: IBeadResolveOptions = { checkShape: false };

      const from = beadCoord.getFromPoint(point, resolveOptions);

      point = point.plus({ x: width, y: height });
      const to = from && beadCoord.getFromPoint(point, resolveOptions);

      from && to
        ? selectionStore.setSelected(createSelection(from, to))
        : selectionStore.setSelected(null);
    } finally {
      selectionStore.toggleSelecting(false);
    }
  }

  function onMouseDown(event: MouseEvent) {
    selectionStore.toggleSelecting(true);
    selectionStore.area.setPoint(event.clientX, event.clientY);
    addEventListener('mouseup', onMouseup, { once: true });
    addEventListener('mousemove', onMouseMove);
  }

  watch(() => toolsStore.isSelection, (isSelectionTool) => {
    if (!isSelectionTool) selectionStore.setSelected(null);
  });

  return computed(() => ({
    mousedown: onMouseDown,
  }));
}
