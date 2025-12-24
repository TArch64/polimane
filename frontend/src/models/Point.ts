export interface IPoint {
  x: number;
  y: number;
}

export class Point implements IPoint {
  static get BLANK(): Point {
    return new Point(0, 0);
  }

  constructor(public x: number, public y: number) {
  }

  plus(other: IPoint | number): Point {
    const x = typeof other === 'number' ? other : other.x;
    const y = typeof other === 'number' ? other : other.y;
    return new Point(this.x + x, this.y + y);
  }

  minus(other: IPoint | number): Point {
    const x = typeof other === 'number' ? other : other.x;
    const y = typeof other === 'number' ? other : other.y;
    return new Point(this.x - x, this.y - y);
  }

  divide(factor: number): Point {
    return new Point(this.x / factor, this.y / factor);
  }

  multiply(factor: number): Point {
    return new Point(this.x * factor, this.y * factor);
  }

  isEqual(other: IPoint): boolean {
    return this.x === other.x && this.y === other.y;
  }

  getAxisDifference(other: IPoint): Point {
    return new Point(
      Math.abs(other.x - this.x),
      Math.abs(other.y - this.y),
    );
  }

  toJSON(): IPoint {
    return {
      x: this.x,
      y: this.y,
    };
  }
}
