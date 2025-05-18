export interface INodeRect {
  x: number;
  y: number;
  width: number;
  height: number;
}

export class NodeRect implements INodeRect {
  static BLANK = new NodeRect({ x: 0, y: 0, width: 0, height: 0 });

  readonly x;
  readonly y;
  readonly width;
  readonly height;

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
      ...patch,
      ...this,
    });
  }

  isEqual(other: INodeRect): boolean {
    return this.x === other.x
      && this.y === other.y
      && this.width === other.width
      && this.height === other.height;
  }
}
