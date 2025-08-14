import type { VNode } from 'vue';
import Konva from 'konva';
import { findKonvaNode } from './findKonvaNode';

export function checkTagAndGetNode(instance: VNode): Konva.Node | null {
  const { el, component } = instance;
  const __konvaNode = findKonvaNode(instance);

  if (el?.tagName && component && !__konvaNode) {
    const name = el.tagName.toLowerCase();
    console.error(
      `vue-konva error: You are trying to render "${name}" inside your component tree. Looks like it is not a Konva node. You can render only Konva components inside the Stage.`,
    );
    return null;
  }

  return __konvaNode;
}
