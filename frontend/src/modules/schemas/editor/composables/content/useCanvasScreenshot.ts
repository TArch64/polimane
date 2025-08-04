import { NodeRect } from '@/models';
import { useEditorStore } from '../../stores';
import { useCanvasStage } from './useCanvasStage';

export function useCanvasScreenshot() {
  const editorStore = useEditorStore();
  const stage = useCanvasStage();

  function generateScreenshot(): string {
    const layer = stage.value.findOne('#editor-layer')!;

    const layerRect = new NodeRect(layer.getClientRect()).delta({
      x: -20,
      y: -10,
      width: 40,
      height: 30,
    });

    return layer.toDataURL({
      ...layerRect.toJSON(),
      mimeType: 'image/webp',
    });
  }

  editorStore.onSaved(async () => {
    await editorStore.updateScreenshot(generateScreenshot());
  });
}
