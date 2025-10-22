export interface IPoint {
  x: number;
  y: number;
}

export class Point implements IPoint {
  static isEqual(a: IPoint, b: IPoint): boolean {
    return a.x === b.x && a.y === b.y;
  }

  readonly x;
  readonly y;

  constructor(point: IPoint) {
    this.x = point.x;
    this.y = point.y;
  }

  get values(): number[] {
    return [this.x, this.y];
  }

  with(patch: Partial<IPoint>): Point {
    return new Point({
      ...this.toJSON(),
      ...patch,
    });
  }

  plus(other: number): Point;
  plus(other: IPoint): Point;
  plus(other: IPoint | number): Point {
    const x = typeof other === 'number' ? other : other.x;
    const y = typeof other === 'number' ? other : other.y;

    return new Point({
      x: this.x + x,
      y: this.y + y,
    });
  }

  div(factor: number): Point {
    return new Point({
      x: this.x / factor,
      y: this.y / factor,
    });
  }

  distanceTo(other: IPoint): number {
    return Math.hypot(this.x - other.x, this.y - other.y);
  }

  toJSON(): IPoint {
    return {
      x: this.x,
      y: this.y,
    };
  }

  isEqual(other: IPoint): boolean {
    return Point.isEqual(this, other);
  }

  getAxisDifference(other: IPoint): IPoint {
    return {
      x: Math.abs(other.x - this.x),
      y: Math.abs(other.y - this.y),
    };
  }
}
