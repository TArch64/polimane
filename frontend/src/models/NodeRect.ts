export interface INodeRect {
  x: number;
  y: number;
  width: number;
  height: number;
}

export class NodeRect implements INodeRect {
  static BLANK = new NodeRect({ x: 0, y: 0, width: 0, height: 0 });

  x;
  y;
  width;
  height;

  constructor(rect: INodeRect) {
    this.x = rect.x;
    this.y = rect.y;
    this.width = rect.width;
    this.height = rect.height;
  }

  get isBlank(): boolean {
    return this.isEqual(NodeRect.BLANK);
  }

  get left(): number {
    return this.x;
  }

  get right(): number {
    return this.x + this.width;
  }

  get top(): number {
    return this.y;
  }

  get bottom(): number {
    return this.y + this.height;
  }

  with(patch: Partial<INodeRect>): NodeRect {
    return new NodeRect({
      ...this,
      ...patch,
    });
  }

  delta(patch: Partial<INodeRect>): NodeRect {
    return new NodeRect({
      x: this.x + (patch.x ?? 0),
      y: this.y + (patch.y ?? 0),
      width: this.width + (patch.width ?? 0),
      height: this.height + (patch.height ?? 0),
    });
  }

  isEqual(other: INodeRect): boolean {
    return this.x === other.x
      && this.y === other.y
      && this.width === other.width
      && this.height === other.height;
  }

  isIntersecting(other: NodeRect): boolean {
    return this.left < other.right
      && this.right > other.left
      && this.top < other.bottom
      && this.bottom > other.top;
  }

  toJSON(): INodeRect {
    return {
      x: this.x,
      y: this.y,
      width: this.width,
      height: this.height,
    };
  }

  clone(): NodeRect {
    return new NodeRect(this.toJSON());
  }
}
