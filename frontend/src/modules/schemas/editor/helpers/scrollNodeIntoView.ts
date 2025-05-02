import Konva from 'konva';

export function scrollNodeIntoView(node: Konva.Node): Promise<void> {
  const stage = node.getStage()!;
  const clientRect = node.getClientRect({ relativeTo: stage });
  const padding = 20;

  const scaleToFit = Math.min(
    (stage.width() - padding * 2) / clientRect.width,
    (stage.height() - padding * 2) / clientRect.height,
  );

  const newScale = Math.min(stage.scaleX(), scaleToFit);
  const targetX = (stage.width() - clientRect.width * newScale) / 2 - clientRect.x * newScale;
  const targetY = (stage.height() - clientRect.height * newScale) / 2 - clientRect.y * newScale;

  return new Promise<void>((resolve) => {
    stage.to({
      x: targetX,
      y: targetY,
      scaleX: newScale,
      scaleY: newScale,
      duration: 0.3,
      easing: Konva.Easings.EaseOut,
      onFinish: resolve,
    });
  });
}
