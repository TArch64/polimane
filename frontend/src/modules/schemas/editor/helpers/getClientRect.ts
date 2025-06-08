import Konva from 'konva';
import { NodeRect } from '@/models';

export function getClientRect(node: Konva.Node): NodeRect {
  const stage = node.getStage();
  const rect = node.getClientRect({ relativeTo: stage! });
  return new NodeRect(rect);
}
