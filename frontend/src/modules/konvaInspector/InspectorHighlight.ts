import Konva from 'konva';

export class InspectorHighlight {
  readonly layer;
  private readonly debugRect;
  private stage!: Konva.Stage;

  constructor() {
    this.debugRect = new Konva.Rect({
      x: 0,
      y: 0,
      width: 0,
      height: 0,
      fill: 'rgba(255, 0, 0, 0.1)',
      stroke: '#000',
      strokeWidth: 2,
    });

    this.layer = new Konva.Layer({
      id: 'konva-inspector-highlight',
      listening: false,
      opacity: 0,
    });
  }

  private isInitialized = false;
  private node: Konva.Node | null = null;

  useStage(stage: Konva.Stage): void {
    if (!this.isInitialized) {
      this.layer.add(this.debugRect);
      this.isInitialized = true;
    } else {
      this.layer.remove();
    }

    stage.add(this.layer);
    this.stage = stage;
  }

  show(node: Konva.Node) {
    if (!this.isHighlightable(node)) {
      return this.hide();
    }

    const previousNode = this.node;

    this.node = node;

    const nodeRect = node.getClientRect({
      relativeTo: this.stage,
    });

    if (!previousNode) {
      this.debugRect.x(nodeRect.x);
      this.debugRect.y(nodeRect.y);
      this.debugRect.width(nodeRect.width);
      this.debugRect.height(nodeRect.height);
      return this.animateLayerOpacity(1);
    }

    this.debugRect.to({
      x: nodeRect.x,
      y: nodeRect.y,
      width: nodeRect.width,
      height: nodeRect.height,
      duration: 0.3,
      easing: Konva.Easings.EaseOut,
    });
  }

  private isHighlightable(node: Konva.Node) {
    return node.nodeType !== 'Stage' && node.nodeType !== 'Layer';
  }

  hide() {
    if (this.node) {
      this.animateLayerOpacity(0);
    }

    this.node = null;
  }

  private animateLayerOpacity(opacity: number) {
    this.layer.to({
      opacity,
      duration: 0.3,
      easing: Konva.Easings.EaseOut,
    });
  }
}
