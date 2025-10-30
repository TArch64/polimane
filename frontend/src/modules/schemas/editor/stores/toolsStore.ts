import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import { EditorTool } from '@editor/enums';
import { type BeadContentKind, BeadKind } from '@/enums';
import { useEditorStore } from './editorStore';

export const useToolsStore = defineStore('schemas/editor/tools', () => {
  const editorStore = useEditorStore();

  const activeTool = ref<EditorTool>(EditorTool.BEAD);
  const isBead = computed(() => activeTool.value === EditorTool.BEAD);
  const isEraser = computed(() => activeTool.value === EditorTool.ERASER);
  const isSelection = computed(() => activeTool.value === EditorTool.SELECTION);
  const isNavigate = computed(() => activeTool.value === EditorTool.NAVIGATE);
  const activateTool = (tool: EditorTool) => activeTool.value = tool;

  const palette = computed({
    get: () => editorStore.schema.palette,
    set: (palette) => editorStore.schema.palette = palette,
  });

  const activeColor = ref<string>(palette.value[0]!);
  const activateColor = (id: string) => activeColor.value = id;

  const activeBead = ref<BeadContentKind>(BeadKind.CIRCLE);
  const activateBead = (bead: BeadContentKind) => activeBead.value = bead;

  return {
    palette,
    activeTool,
    isBead,
    isEraser,
    isSelection,
    isNavigate,
    activateTool,
    activeColor,
    activateColor,
    activeBead,
    activateBead,
  };
});
