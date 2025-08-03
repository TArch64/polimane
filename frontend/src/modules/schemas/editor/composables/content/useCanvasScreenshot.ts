import { useEditorStore } from '../../stores';
import { useCanvasStage } from './useCanvasStage';

export function useCanvasScreenshot() {
  const editorStore = useEditorStore();
  const stage = useCanvasStage();

  editorStore.onSaved(() => {

  });
}
