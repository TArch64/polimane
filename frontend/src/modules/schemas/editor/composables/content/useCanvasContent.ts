import { toRef } from 'vue';
import { useEditorStore } from '../../stores';
import { onCanvasReady } from '../onCanvasReady';
import { initObjectRegistry } from './useObjectRegistry';
import { usePatternRenderer } from './pattern';

export function useCanvasContent() {
  const editorStore = useEditorStore();
  initObjectRegistry();

  onCanvasReady(() => {
    usePatternRenderer(toRef(editorStore.schema, 'content'));
  });
}
