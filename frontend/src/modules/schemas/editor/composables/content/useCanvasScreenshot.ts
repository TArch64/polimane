import Konva from 'konva';
import { NodeRect } from '@/models';
import { useEditorStore } from '../../stores';
import { useCanvasStage } from './useCanvasStage';

export const SCREENSHOT_IGNORE = 'ignore-on-screenshot';

export function useCanvasScreenshot() {
  const editorStore = useEditorStore();
  const stage = useCanvasStage();

  function buildScreenshotLayer(): Konva.Layer {
    const original = stage.value.findOne('#editor-layer')!;
    const layer: Konva.Layer = original.clone({ listening: false });

    const ignoredNodes = layer.find(`.${SCREENSHOT_IGNORE}`);

    for (const ignoredNode of ignoredNodes) {
      ignoredNode.remove();
    }

    return layer;
  }

  function generateScreenshot(): string {
    const layer = buildScreenshotLayer();

    const layerRect = new NodeRect(layer.getClientRect()).delta({
      x: -20,
      y: -20,
      width: 40,
      height: 40,
    });

    return layer.toDataURL({
      ...layerRect.toJSON(),
      mimeType: 'image/webp',
      pixelRatio: window.devicePixelRatio,
    });
  }

  function needScreenshot(): boolean {
    if (!stage.value) {
      return false;
    }

    if (!editorStore.schema.screenshotedAt) {
      return true;
    }

    const lastSaved = new Date(editorStore.schema.screenshotedAt);
    const now = new Date();
    const diff = now.getTime() - lastSaved.getTime();

    return diff > 30 * 60 * 1000;
  }

  editorStore.onSaved(async () => {
    if (needScreenshot()) {
      await editorStore.updateScreenshot(generateScreenshot());
    }
  });
}
