import Konva from 'konva';
import { NodeRect } from '@/models';
import { useEditorStore } from '../stores';
import { useCanvasStage } from './content';

export const SCREENSHOT_IGNORE = 'ignore-on-screenshot';

export function useEditorScreenshot() {
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

  function generateScreenshot(): string | null {
    const layer = buildScreenshotLayer();
    const contentRect = new NodeRect(layer.getClientRect());

    if (contentRect.isBlank) {
      return null;
    }

    const targetRect = contentRect.delta({
      x: -20,
      y: -20,
      width: 40,
      height: 40,
    });

    return layer.toDataURL({
      ...targetRect.toJSON(),
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
    const source = needScreenshot() ? generateScreenshot() : null;
    if (source) await editorStore.updateScreenshot(source);
  });
}
