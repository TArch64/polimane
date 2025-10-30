import type { Ref } from 'vue';
import type { IBeadsGrid } from '@editor/composables';

type ToolListenerEvent = 'mousedown' | 'mousemove' | 'mouseup' | 'click';
type ToolListener<E extends Event> = (event: E) => void;

export type EditorToolListeners = Partial<{
  [K in ToolListenerEvent]: ToolListener<SVGElementEventMap[K]>;
}>;

export interface IEditorToolOptions {
  canvasRef: Ref<SVGSVGElement>;
  groupRef: Ref<SVGGElement>;
  backgroundRef: Ref<SVGRectElement>;
  beadsGrid: IBeadsGrid;
}

export interface IEditorTool {
  level: 'canvas' | 'content';
  onActivated?: (abortController: AbortController) => void;
  onDeactivated?: () => void;
  listeners: EditorToolListeners;
}
