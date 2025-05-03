import Konva from 'konva';

export interface IScrollNodeIntoViewOptions {
  scale?: boolean;
}

export function scrollNodeIntoView(
  node: Konva.Node,
  options: IScrollNodeIntoViewOptions = {},
): Promise<void> {
  const stage = node.getStage()!;
  const clientRect = node.getClientRect({ relativeTo: stage });
  const padding = 20;

  // Determine whether to scale or not (defaults to true if not specified)
  const shouldScale = options.scale !== false;

  // Calculate scale only if scaling is enabled
  let newScale = stage.scaleX();
  if (shouldScale) {
    // Calculate scale to fit node with padding (will scale up or down as needed)
    newScale = Math.min(
      (stage.width() - padding * 2) / clientRect.width,
      (stage.height() - padding * 2) / clientRect.height,
    );
  }

  // Calculate position to center the node
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
