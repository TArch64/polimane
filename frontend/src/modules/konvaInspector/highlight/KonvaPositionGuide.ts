import Konva from 'konva';
import { Point } from '@/models';

export interface IKonvaPositionGuidePosition {
  from: Point;
  to: Point;
}

export interface IKonvaPositionGuideProps {
  position: IKonvaPositionGuidePosition;
}

export class KonvaPositionGuide extends Konva.Label {
  private readonly relativeFrom: Point;
  private readonly relativeTo: Point;

  constructor(props: IKonvaPositionGuideProps) {
    const [from, to] = props.position.from.x < props.position.to.x || props.position.from.y < props.position.to.y
      ? [props.position.from, props.position.to]
      : [props.position.to, props.position.from];

    const relativeTo = new Point({
      x: (to.x - from.x) / 2,
      y: (to.y - from.y) / 2,
    });

    super({
      x: from.x + relativeTo.x,
      y: from.y + relativeTo.y,
    });

    this.relativeFrom = new Point({
      x: -relativeTo.x,
      y: -relativeTo.y,
    });

    this.relativeTo = relativeTo;

    const tag = this.createTag();

    const text = this.createText();
    this.offsetX(text.width() / 2);
    this.offsetY(text.height() / 2);

    const line = this.createLine();
    this.add(line, tag, text);
  }

  private createLine(): Konva.Line {
    return new Konva.Line({
      points: [
        ...this.relativeFrom.plus(this.offset()).values,
        ...this.relativeTo.plus(this.offset()).values,
      ],
      stroke: '#000',
      strokeWidth: 1,
    });
  }

  private createTag(): Konva.Tag {
    return new Konva.Tag({
      ...this.relativeTo.div(2).toJSON(),
      fill: '#fff',
      stroke: '#000',
      strokeWidth: 1,
      cornerRadius: 4,
    });
  }

  private createText(): Konva.Text {
    return new Konva.Text({
      text: `${this.relativeFrom.distanceTo(this.relativeTo)}px`,
      fontSize: 12,
      fill: '#000',
      padding: 5,
    });
  }
}
