import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import { useEditorStore } from './editorStore';

export type ActiveToolId = 'eraser' | number;

export const useToolsStore = defineStore('schemas/editor/tools', () => {
  const editorStore = useEditorStore();

  const palette = computed({
    get: () => editorStore.schema.palette,
    set: (palette) => editorStore.schema.palette = palette,
  });

  const activeToolId = ref<ActiveToolId>(0);
  const activateTool = (id: ActiveToolId) => activeToolId.value = id;

  const activeColorId = computed((): number => {
    return typeof activeToolId.value !== 'number' ? -1 : activeToolId.value;
  });

  const activeColor = computed(() => {
    return palette.value[activeColorId.value] || '';
  });

  return {
    palette,
    activeColor,
    activeColorId,
    activeToolId,
    activateTool,
  };
});
