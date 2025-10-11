import { computed, type Ref, watch } from 'vue';
import {
  type IBeadSelection,
  useBeadsStore,
  useCanvasStore,
  useSelectionStore,
  useToolsStore,
} from '@editor/stores';
import {
  parseSchemaBeadCoord,
  type SchemaBeadCoord,
  type SchemaBeadCoordTuple,
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
  const canvasStore = useCanvasStore();

  const beadCoord = useBeadCoord(options);

  function createSelection(from: SchemaBeadCoordTuple, to: SchemaBeadCoordTuple): IBeadSelection | null {
    const selected = Object.keys(beadsStore.getInArea(from, to)) as SchemaBeadCoord[];

    if (!selected.length) {
      return null;
    }

    if (selected.length === 1) {
      return { from: selected[0]!, to: selected[0]! };
    }

    const parsed = selected.map(parseSchemaBeadCoord);
    const xs = parsed.map(([x]) => x);
    const ys = parsed.map(([, y]) => y);

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
