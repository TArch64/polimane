import type { Ref } from 'vue';
import type { IBeadsGrid } from '@editor/composables';

type ToolListenerEvent = 'mousedown' | 'mousemove' | 'mouseup';
type ToolListener = (event: MouseEvent) => void;
export type ToolListeners = Partial<Record<ToolListenerEvent, ToolListener>>;

export interface IEditorTool {
  level: 'canvas' | 'content';
  listeners: ToolListeners;
}

export interface IToolsOptions {
  canvasRef: Ref<SVGSVGElement>;
  groupRef: Ref<SVGGElement>;
  backgroundRef: Ref<SVGRectElement>;
  beadsGrid: IBeadsGrid;
}

export type UseEditorTool = (options: IToolsOptions) => Ref<IEditorTool>;
