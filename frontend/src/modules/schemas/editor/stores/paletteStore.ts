import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import { useEditorStore } from './editorStore';

export type ActiveToolId = 'eraser' | number;

export const usePaletteStore = defineStore('schemas/editor/palette', () => {
  const editorStore = useEditorStore();

  const palette = computed({
    get: () => editorStore.schema.palette,
    set: (colors) => editorStore.schema.palette = colors,
  });

  const activeToolId = ref<ActiveToolId>(0);
  const activateTool = (id: ActiveToolId) => activeToolId.value = id;

  const activeColor = computed(() => {
    return activeToolId.value === 'eraser' ? '' : palette.value[activeToolId.value] || '';
  });

  return {
    palette,
    activeColor,
    activeToolId,
    activateTool,
  };
});
