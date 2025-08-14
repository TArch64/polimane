import type { VNode } from 'vue';
import Konva from 'konva';

export function findKonvaNode(instance: VNode): Konva.Node | null {
  if (!instance.component) return null;

  return instance.component.__konvaNode || findKonvaNode(instance.component.subTree);
}
