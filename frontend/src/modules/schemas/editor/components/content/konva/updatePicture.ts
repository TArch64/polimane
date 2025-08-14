import Konva from 'konva';

export function updatePicture(node: Konva.Node) {
  if (!Konva.autoDrawEnabled) {
    const drawingNode = node.getLayer() || node.getStage();
    drawingNode && drawingNode.batchDraw();
  }
}
