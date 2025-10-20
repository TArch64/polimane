import { computed, type Ref, watch } from 'vue';
import {
  type IBeadSelection,
  useBeadsStore,
  useSelectionStore,
  useToolsStore,
} from '@editor/stores';
import { BEAD_SIZE } from '@editor/const';
import { type IPoint, parseBeadCoord, Point, serializeBeadCoord } from '@/models';
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

    for (const coord of selected.map(parseBeadCoord)) {
      xs.push(coord.x);
      ys.push(coord.y);
    }

    return {
      from: serializeBeadCoord(Math.min(...xs), Math.min(...ys)),
      to: serializeBeadCoord(Math.max(...xs), Math.max(...ys)),
    };
  }

  function onMouseMove(event: MouseEvent) {
    selectionStore.area.extend(event.movementX, event.movementY);
  }

  function setCanvasSelection() {
    const { from, to } = selectionStore.selected!;
    const fromOffset = options.beadsGrid.resolveBeadOffset(from);
    const toOffset = options.beadsGrid.resolveBeadOffset(to);

    const PADDING = 2;
    const x = Math.min(fromOffset.x, toOffset.x) - PADDING;
    const y = Math.min(fromOffset.y, toOffset.y) - PADDING;
    const width = Math.abs(fromOffset.x - toOffset.x) + BEAD_SIZE + PADDING * 2;
    const height = Math.abs(fromOffset.y - toOffset.y) + BEAD_SIZE + PADDING * 2;

    selectionStore.area.setPoint(x, y);
    selectionStore.area.extend(width, height);
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

      selectionStore.selected
        ? setCanvasSelection()
        : selectionStore.reset();
    } finally {
      selectionStore.toggleSelecting(false);
    }
  }

  function onMouseDown(event: MouseEvent) {
    selectionStore.setSelected(null);
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
