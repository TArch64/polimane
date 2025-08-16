import type { CustomInspectorNode } from '@vue/devtools-kit';
import Konva from 'konva';

export class InspectorTree {
  constructor(private readonly excludeId: string) {
  }

  collectInspectorTree(stage: Konva.Stage): CustomInspectorNode[] {
    return [this.collectNodeTree(stage)];
  }

  private collectNodeTree(node: Konva.Node): CustomInspectorNode {
    const inspectorNode: CustomInspectorNode = {
      id: node._id.toString(),
      label: this.getNodeLabel(node),
    };

    if (node instanceof Konva.Container) {
      inspectorNode.children = node.children
        .filter((child) => child.id() !== this.excludeId)
        .map((child) => this.collectNodeTree(child));
    }

    return inspectorNode;
  }

  private getNodeLabel(node: Konva.Node): string {
    if (node.id()) return `#${node.id()}`;
    if (node.name()) return `.${node.name()}`;
    return node.className || node.nodeType;
  }
}
