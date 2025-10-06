import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import { useEditorStore } from './editorStore';

export type ActiveToolId = 'eraser' | 'bead';

export const useToolsStore = defineStore('schemas/editor/tools', () => {
  const editorStore = useEditorStore();

  const palette = computed({
    get: () => editorStore.schema.palette,
    set: (palette) => editorStore.schema.palette = palette,
  });

  const activeTool = ref<ActiveToolId>('bead');
  const isEraser = computed(() => activeTool.value === 'eraser');
  const activateTool = (id: ActiveToolId) => activeTool.value = id;

  const activeColor = ref<string>(palette.value[0]!);
  const activateColor = (id: string) => activeColor.value = id;

  return {
    palette,
    activeTool,
    isEraser,
    activateTool,
    activeColor,
    activateColor,
  };
});
