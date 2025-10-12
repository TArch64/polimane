import { computed, type Ref, watch } from 'vue';
import {
  type IBeadSelection,
  useBeadsStore,
  useSelectionStore,
  useToolsStore,
} from '@editor/stores';
import {
  type IPoint,
  parseSchemaBeadCoord,
  type SchemaBeadCoord,
  serializeSchemaBeadCoord,
} from '@/models';
import type { IBeadToolsOptions } from './IBeadToolsOptions';
import { useBeadCoord } from './useBeadCoord';

export interface IBeadSelectionListeners {
  mousedown: (event: MouseEvent) => void;
}

export function useBeadSelection(options: IBeadToolsOptions): Ref<IBeadSelectionListeners> {
  const selectionStore = useSelectionStore();
  const toolsStore = useToolsStore();
  const beadsStore = useBeadsStore();

  const beadCoord = useBeadCoord(options);

  function createSelection(from: IPoint, to: IPoint): IBeadSelection | null {
    const selected = Object.keys(beadsStore.getInArea(from, to)) as SchemaBeadCoord[];

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
    selectionStore.extend(event.movementX, event.movementY);
  }

  function onMouseup() {
    removeEventListener('mousemove', onMouseMove);
    beadCoord.clearCache();

    const baseX = selectionStore.selection.x;
    const baseY = selectionStore.selection.y;

    const from = beadCoord.getFromPoint({ x: baseX, y: baseY }, {
      checkShape: false,
    });

    const to = from && beadCoord.getFromPoint({
      x: baseX + selectionStore.selection.width,
      y: baseY + selectionStore.selection.height,
    }, {
      checkShape: false,
    });

    selectionStore.toggleSelecting(false);

    from && to
      ? selectionStore.setSelected(createSelection(from, to))
      : selectionStore.setSelected(null);
  }

  function onMouseDown(event: MouseEvent) {
    selectionStore.toggleSelecting(true);
    selectionStore.setPoint(event.clientX, event.clientY);
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
