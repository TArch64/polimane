export interface IPoint {
  x: number;
  y: number;
}

export class Point implements IPoint {
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
    return Math.sqrt(Math.pow(this.x - other.x, 2) + Math.pow(this.y - other.y, 2));
  }

  toJSON(): IPoint {
    return {
      x: this.x,
      y: this.y,
    };
  }

  isEqual(other: IPoint): boolean {
    return this.x === other.x && this.y === other.y;
  }
}
