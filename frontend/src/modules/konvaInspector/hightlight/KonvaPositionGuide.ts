import Konva from 'konva';

export interface IKonvaPositionGuidePosition {
  from: Konva.Vector2d;
  to: Konva.Vector2d;
}

export interface IKonvaPositionGuideProps {
  position: IKonvaPositionGuidePosition;
}

export class KonvaPositionGuide extends Konva.Label {
  private readonly relativeFrom: Konva.Vector2d;
  private readonly relativeTo: Konva.Vector2d;

  constructor(props: IKonvaPositionGuideProps) {
    const [from, to] = props.position.from.x < props.position.to.x || props.position.from.y < props.position.to.y
      ? [props.position.from, props.position.to]
      : [props.position.to, props.position.from];

    const relativeTo = {
      x: (to.x - from.x) / 2,
      y: (to.y - from.y) / 2,
    };

    super({
      x: from.x + relativeTo.x,
      y: from.y + relativeTo.y,
    });

    this.relativeFrom = {
      x: -relativeTo.x,
      y: -relativeTo.y,
    };

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
        this.relativeFrom.x + this.offsetX(),
        this.relativeFrom.y + this.offsetY(),
        this.relativeTo.x + this.offsetX(),
        this.relativeTo.y + this.offsetY(),
      ],
      stroke: '#000',
      strokeWidth: 1,
    });
  }

  private createTag(): Konva.Tag {
    return new Konva.Tag({
      x: this.relativeTo.x / 2,
      y: this.relativeTo.y / 2,
      fill: '#fff',
      stroke: '#000',
      strokeWidth: 1,
      cornerRadius: 4,
    });
  }

  private createText(): Konva.Text {
    const distance = Math.sqrt(Math.pow(this.relativeTo.x, 2) + Math.pow(this.relativeTo.y, 2)).toString();
    return new Konva.Text({
      text: `${distance}px`,
      fontSize: 12,
      fill: '#000',
      padding: 5,
    });
  }
}
