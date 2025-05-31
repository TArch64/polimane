import { defineStore } from 'pinia';
import { computed } from 'vue';
import { useEditorStore } from './editorStore';

export const usePaletteStore = defineStore('schemas/editor/palette', () => {
  const editorStore = useEditorStore();

  const palette = computed({
    get: () => editorStore.schema.palette,
    set: (colors) => editorStore.schema.palette = colors,
  });

  return { palette };
});
