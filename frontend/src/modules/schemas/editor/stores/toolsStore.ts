import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import { EditorTool } from '@editor/enums';
import { BeadKind } from '@/enums';
import { useEditorStore } from './editorStore';

export const useToolsStore = defineStore('schemas/editor/tools', () => {
  const editorStore = useEditorStore();

  const activeTool = ref<EditorTool>(EditorTool.BEAD);
  const isBead = computed(() => activeTool.value === EditorTool.BEAD);
  const isEraser = computed(() => activeTool.value === EditorTool.ERASER);
  const isSelection = computed(() => activeTool.value === EditorTool.SELECTION);
  const activateTool = (tool: EditorTool) => activeTool.value = tool;

  const palette = computed({
    get: () => editorStore.schema.palette,
    set: (palette) => editorStore.schema.palette = palette,
  });

  const activeColor = ref<string>(palette.value[0]!);
  const activateColor = (id: string) => activeColor.value = id;

  const activeBead = ref<BeadKind>(BeadKind.CIRCLE);
  const activateBead = (bead: BeadKind) => activeBead.value = bead;

  return {
    palette,
    activeTool,
    isBead,
    isEraser,
    isSelection,
    activateTool,
    activeColor,
    activateColor,
    activeBead,
    activateBead,
  };
});
