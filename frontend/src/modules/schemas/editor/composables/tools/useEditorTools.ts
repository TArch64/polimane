import { computed, onMounted, watch } from 'vue';
import { useToolsStore } from '@editor/stores';
import { getObjectKeys } from '@/helpers';
import type { IEditorTool, IToolsOptions } from './tool';
import { usePaintingTool } from './usePaintingTool';
import { useSelectionTool } from './useSelectionTool';
import { useNavigateTool } from './useNavigateTool';

export function useEditorTools(options: IToolsOptions) {
  const toolsStore = useToolsStore();
  const painting = usePaintingTool(options);
  const selection = useSelectionTool(options);
  const navigate = useNavigateTool(options);

  const tool = computed(() => {
    if (toolsStore.isSelection) return selection.value;
    if (toolsStore.isNavigate) return navigate.value;
    return painting.value;
  });

  function getTarget(tool: IEditorTool): SVGElement {
    const { canvasRef, groupRef } = options;
    return tool.level === 'canvas' ? canvasRef.value : groupRef.value;
  }

  let abortController: AbortController | null = null;

  function updateTool(newTool: IEditorTool) {
    abortController?.abort();
    abortController = new AbortController();

    const target = getTarget(newTool);

    for (const event of getObjectKeys(newTool.listeners)) {
      const listener = newTool.listeners[event];
      if (listener) {
        target.addEventListener(event, listener, {
          signal: abortController.signal,
        });
      }
    }
  }

  onMounted(() => updateTool(tool.value));
  watch(tool, updateTool, { deep: true });
}
