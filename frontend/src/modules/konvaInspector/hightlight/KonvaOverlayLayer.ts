import Konva from 'konva';
import { NodeRect, Point } from '@/models';
import { type IKonvaPositionGuidePosition, KonvaPositionGuide } from './KonvaPositionGuide';

export interface IKonvaOverlayLayerProps {
  id: string;
  parentRect: NodeRect;
  targetRect: NodeRect;
}

export class KonvaOverlayLayer extends Konva.Layer {
  constructor(props: IKonvaOverlayLayerProps) {
    super({
      id: props.id,
      listening: false,
    });

    this.add(this.createParentHighlight(props.parentRect));
    this.add(this.createTargetHighlight(props.targetRect));

    this.add(this.createLeftPositionGuide(props.parentRect, props.targetRect));
    this.add(this.createTopPositionGuide(props.parentRect, props.targetRect));
    this.add(this.createRightPositionGuide(props.parentRect, props.targetRect));
    this.add(this.createBottomPositionGuide(props.parentRect, props.targetRect));
  }

  private createParentHighlight(rect: NodeRect): Konva.Rect {
    return new Konva.Rect({
      x: rect.x,
      y: rect.y,
      width: rect.width,
      height: rect.height,
      stroke: '#000',
      strokeWidth: 1,
    });
  }

  private createTargetHighlight(rect: NodeRect): Konva.Rect {
    return new Konva.Rect({
      x: rect.x,
      y: rect.y,
      width: rect.width,
      height: rect.height,
      fill: 'rgba(255, 0, 0, 0.1)',
      stroke: 'rgba(0, 0, 0, 0.5)',
      strokeWidth: 1,
    });
  }

  private createLeftPositionGuide(parentRect: NodeRect, targetRect: NodeRect): KonvaPositionGuide {
    const y = targetRect.y + targetRect.height / 2;

    return this.createPositionGuide({
      from: new Point({ x: parentRect.left, y }),
      to: new Point({ x: targetRect.left, y }),
    });
  }

  private createTopPositionGuide(parentRect: NodeRect, targetRect: NodeRect): KonvaPositionGuide {
    const x = targetRect.x + targetRect.width / 2;

    return this.createPositionGuide({
      from: new Point({ x, y: parentRect.top }),
      to: new Point({ x, y: targetRect.top }),
    });
  }

  private createRightPositionGuide(parentRect: NodeRect, targetRect: NodeRect): KonvaPositionGuide {
    const y = targetRect.y + targetRect.height / 2;

    return this.createPositionGuide({
      from: new Point({ x: parentRect.right, y }),
      to: new Point({ x: targetRect.right, y }),
    });
  }

  private createBottomPositionGuide(parentRect: NodeRect, targetRect: NodeRect): KonvaPositionGuide {
    const x = targetRect.x + targetRect.width / 2;

    return this.createPositionGuide({
      from: new Point({ x, y: parentRect.bottom }),
      to: new Point({ x, y: targetRect.bottom }),
    });
  }

  private createPositionGuide(position: IKonvaPositionGuidePosition): KonvaPositionGuide {
    return new KonvaPositionGuide({ position });
  }
}
