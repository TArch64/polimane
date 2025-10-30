import { computed, onMounted, watch } from 'vue';
import { useToolsStore } from '@editor/stores';
import { getObjectKeys } from '@/helpers';
import type { SafeAny } from '@/types';
import type { IEditorTool, IEditorToolOptions } from './tool';
import { usePaintingTool } from './usePaintingTool';
import { useSelectionTool } from './useSelectionTool';
import { useNavigateTool } from './useNavigateTool';
import { useZoomTool } from './useZoomTool';

export function useEditorTools(options: IEditorToolOptions) {
  const toolsStore = useToolsStore();
  const painting = usePaintingTool(options);
  const selection = useSelectionTool(options);
  const navigate = useNavigateTool();
  const zoom = useZoomTool();

  const tool = computed(() => {
    if (toolsStore.isSelection) return selection.value;
    if (toolsStore.isNavigate) return navigate.value;
    if (toolsStore.isZoom) return zoom.value;
    return painting.value;
  });

  function getTarget(tool: IEditorTool): SVGElement {
    const { canvasRef, groupRef } = options;
    return tool.level === 'canvas' ? canvasRef.value : groupRef.value;
  }

  let abortController: AbortController | null = null;

  function updateTool(newTool: IEditorTool, oldTool?: IEditorTool) {
    abortController?.abort();
    oldTool?.onDeactivated?.();
    abortController = new AbortController();

    const target = getTarget(newTool);

    for (const event of getObjectKeys(newTool.listeners)) {
      const listener = newTool.listeners[event];
      if (listener) {
        target.addEventListener(event, listener as SafeAny, {
          signal: abortController.signal,
        });
      }
    }

    newTool.onActivated?.(abortController);
  }

  onMounted(() => updateTool(tool.value));
  watch(tool, updateTool, { deep: true });
}
