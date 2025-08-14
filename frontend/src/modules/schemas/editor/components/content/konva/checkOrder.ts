import type { VNode } from 'vue';
import Konva from 'konva';
import { getChildren } from './getChildren';
import { checkTagAndGetNode } from './checkTagAndGetNode';
import { updatePicture } from './updatePicture';

export function checkOrder(subTree: VNode, konvaNode: Konva.Node) {
  const children = getChildren(subTree);

  const nodes: Konva.Node[] = [];
  children.forEach((child) => {
    const konvaNode = checkTagAndGetNode(child);
    if (konvaNode) {
      nodes.push(konvaNode);
    }
  });

  let needRedraw = false;
  nodes.forEach((konvaNode, index) => {
    if (konvaNode.getZIndex() !== index) {
      konvaNode.setZIndex(index);
      needRedraw = true;
    }
  });

  if (needRedraw) {
    updatePicture(konvaNode);
  }
}
