import Konva from 'konva';
import { NodeRect } from '@/models';
import { KonvaOverlayLayer } from './KonvaOverlayLayer';

export class InspectorHighlight {
  readonly layerId = 'konva-inspector-highlight';
  private stage!: Konva.Stage;
  private layer: KonvaOverlayLayer | null = null;

  useStage(stage: Konva.Stage): void {
    this.stage = stage;
  }

  show(node: Konva.Node) {
    if (!this.isHighlightable(node)) {
      return this.hide();
    }

    this.layer?.destroy();

    this.layer = new KonvaOverlayLayer({
      id: this.layerId,
      parentRect: this.getNodeRect(node.parent!),
      targetRect: this.getNodeRect(node),
    });

    this.stage.add(this.layer);
  }

  hide() {
    this.layer?.remove();
    this.layer = null;
  }

  private isHighlightable(node: Konva.Node) {
    return node.nodeType !== 'Stage' && node.nodeType !== 'Layer';
  }

  private getNodeRect(node: Konva.Node): NodeRect {
    const json = node.getClientRect({
      relativeTo: this.stage,
    });

    return new NodeRect(json);
  }
}
