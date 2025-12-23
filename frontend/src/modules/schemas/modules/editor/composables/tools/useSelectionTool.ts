import { computed, watch } from 'vue';
import {
  type IBeadSelection,
  useBeadsStore,
  useSelectionStore,
  useToolsStore,
} from '@editor/stores';
import { BEAD_SIZE } from '@editor/const';
import { parseBeadCoord, Point, serializeBeadCoord } from '@/models';
import { getObjectKeys } from '@/helpers';
import type { IEditorTool, IEditorToolOptions } from './tool';
import { type IBeadResolveOptions, useBeadCoord } from './useBeadCoord';

export const useSelectionTool = (options: IEditorToolOptions) => {
  const selectionStore = useSelectionStore();
  const toolsStore = useToolsStore();
  const beadsStore = useBeadsStore();

  const beadCoord = useBeadCoord(options);

  function createSelection(from: Point, to: Point): IBeadSelection | null {
    const selected = beadsStore.getInArea(from, to);

    if (!Object.keys(selected).length) {
      return null;
    }

    const xs: number[] = [];
    const ys: number[] = [];

    for (const coord of getObjectKeys(selected)) {
      const { x, y } = parseBeadCoord(coord);
      xs.push(x);
      ys.push(y);
    }

    xs.sort((a, b) => a - b);
    ys.sort((a, b) => a - b);

    return {
      from: serializeBeadCoord(xs[0]!, ys[0]!),
      to: serializeBeadCoord(xs.at(-1)!, ys.at(-1)!),
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
      let point = new Point(x, y);
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

  return computed((): IEditorTool => ({
    level: 'content',
    listeners: { mousedown: onMouseDown },
  }));
};
